# SES-Journal — архитектура и заметки

Конспект разбора архитектуры проекта и рефакторинга обработки ошибок / внедрения зависимостей.

## Общая идея

Проект построен по слоистой архитектуре (Layered / Clean Architecture):

```
handler → service → repository → gorm/Postgres
```

Каждый слой зависит только от слоя ниже и не знает о деталях соседей.

| Слой | Файл | Ответственность | Чего не должен знать |
|---|---|---|---|
| Router | `internal/router/router.go` | Привязка URL к обработчикам | про SQL и бизнес-логику |
| Handler | `internal/handler/*.go` | HTTP: разбор запроса, сериализация ответа | про БД и GORM |
| Service | `internal/services/userService.go` | Бизнес-логика | про HTTP и хранение данных |
| Repository | `internal/repository/usersRepo.go` | Доступ к данным (GORM/SQL) | про HTTP и бизнес-правила |
| DB | `internal/repository/db.go` | Соединение с Postgres | — |

Выгода разделения: смена БД/ORM затрагивает только `repository`, смена транспорта — только `handler`, сервис тестируется с мок-репозиторием.

## Поток запуска и обработки запроса

```
main.go
  └─ app.Run()
       ├─ conf.GetAppConfig(), logger.Init()
       ├─ repository.GetDBConnection()      → синглтон *gorm.DB
       ├─ repository.NewUsersRepo(db)        → UsersRepository
       ├─ services.NewUserService(urepo)     → *UserService
       ├─ handler.NewUserHandler(uservice)   → *UserHandler
       ├─ router.New() / rt.Register(uhandler) → http.Handler
       └─ http.Server.ListenAndServe()

HTTP GET /users
  └─ UserHandler.HandleUsersList()          (транспортный слой)
       ├─ h.UserService.GetUsersList()
       │    └─ repo.GetUsersList()          → GORM-запрос в Postgres
       └─ json.Marshal(result) → w.Write()
```

### Важный нюанс: когда выполняется тело функции

Тело функции не выполняется в момент её определения или упоминания — только при фактическом вызове `()`.

- `router.New()` создаёт только `Router{mux}`; `NewHealthHandler`/`NewUserHandler` **не** вызываются.
- `rt.Register(...)` вызывает тело метода: создаёт хендлеры, вешает роуты, возвращает `mux`.
- `HandleUsersList` вызывается на каждый HTTP-запрос.

## Внедрение зависимостей (DI)

Зависимости собираются один раз при старте — в композиционном корне `app.Run()`, и передаются вниз через конструкторы `NewXxx`:

```go
db       := repository.GetDBConnection()
urepo    := repository.NewUsersRepo(db)
uservice := services.NewUserService(urepo)
uhandler := handler.NewUserHandler(uservice)

rt := router.New()
h  := rt.Register(uhandler)
```

`repository.GetDBConnection()` — ленивый глобальный синглтон: соединение создаётся при первом обращении и переиспользуется.

Репозиторий скрыт за интерфейсом (`repository.UsersRepository`), поэтому сервис зависит от абстракции, а не от конкретного типа — это позволяет подменять реализацию в тестах.

## Обработка ошибок: почему `Fatal` опасен внутри запроса

`logrus.Fatal` пишет лог и вызывает `os.Exit(1)` — процесс завершается немедленно. Это уместно **только на старте** приложения, когда без ошибки оно всё равно неработоспособно:

- `internal/config/app.go` — загрузка `.env`;
- `internal/repository/db.go` — подключение, миграции, сиды;
- `internal/app/app.go` — падение `ListenAndServe`.

Внутри обработки запроса `Fatal` недопустим: одна неудачная выборка или битый JSON обрушит сервер для всех.

**Принцип:** ошибку возвращают вверх по стеку и обрабатывают на границе (в хендлере, где известно, что это HTTP). Логируют один раз — там, где ошибка обработана.

### Что было и что стало

| Место | Было | Стало |
|---|---|---|
| `services/userService.go` | `Fatal` + глотание ошибки (`return nil, nil`) | проброс ошибки наверх (`return s.repo.GetUsersList()`) |
| `handler/userHandler.go` (ошибка сервиса) | `Fatal` | `Error`-лог + `http.Error(..., 500)` + `return` |
| `handler/userHandler.go` (ошибка marshal) | `Fatal` | `Error`-лог + `http.Error(..., 500)` + `return` |
| Стартовые пути | `Fatal` | оставлено (`Fatal` оправдан) |

Ключевой баг старой версии хендлера: после `http.Error` не было `return` — выполнение продолжалось, `nil`-результат сериализовался в `"null"`, а `w.Write` дописывал тело к уже отправленному ответу 500.

## Итоговое состояние после рефакторинга

- DI-цепочка строится в `app.go:30-37`; хендлер получает сервис, сервис — репозиторий, репозиторий — `*gorm.DB`.
- `UsersRepository` (интерфейс) с корректной сигнатурой `GetUsersList() ([]models.User, error)`; `NewUsersRepo` возвращает интерфейс, сервис зависит от интерфейса.
- `Fatal` остался только в стартовых путях.
- Мёртвый закомментированный код удалён.

### Возможные улучшения на будущее

- Назвать интерфейс `UserRepo` (короче, без дублирования с `UsersRepo`).
- Поле `UserHandler.UserService` сделать неэкспортируемым (`userService`).
- Обрабатывать ошибку `w.Write`.
- Отдавать JSON-тело ошибки вместо `text/plain` от `http.Error`.
- Graceful shutdown через `server.Shutdown` и обработку сигналов вместо `Fatal` на `ListenAndServe`.
