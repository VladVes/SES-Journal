package handler

import (
	"encoding/json"

	"net/http"

	"github.com/VladVes/SES-Journal/internal/dto"
	"github.com/VladVes/SES-Journal/internal/logger"
	"github.com/VladVes/SES-Journal/internal/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userv *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userv,
	}
}

func (h *UserHandler) HandleUsersList(w http.ResponseWriter, r *http.Request) {
	result, err := h.userService.GetUsersList()
	if err != nil {
		logger.Log.WithError(err).Error("error while getting users list")
		http.Error(w, "internal server error while getting users list", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		logger.Log.WithError(err).Error("error while users list json serialize")
		http.Error(w, "internal server error while users list json serialize", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func (h *UserHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// defer r.Body.Close() — хорошая привычка
	defer r.Body.Close()

	// Чтобы защититься от огромных тел, можно обернуть:
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // (1МБ) перед декодированием
	// 1<<20 побитовый сдвиг 0b00000001 - сдвигаем 1 на 20 позиций влево что эквивалентно 2^20 т.е. 1048576bit == 1Mb

	// json.NewDecoder(r.Body) — лучше, чем io.ReadAll + json.Unmarshal, меньше аллокаций, читает прямо из тела запроса
	dec := json.NewDecoder(r.Body)

	// dec.DisallowUnknownFields() — полезно, чтобы сразу отловить неверный JSON
	dec.DisallowUnknownFields()

	var req dto.CreateUserRequest
	if err := dec.Decode(&req); err != nil {
		// для 400тых можно Warn
		logger.Log.WithError(err).Warn("invalid json body for create user")
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// нормализуем и валидируем методами dto
	// если номрализация и валидация как-то связаны с бизнес-логикой
	// то лучше это делать на уровне sevice
	req.Normalize()
	if err := req.Validate(); err != nil {
		logger.Log.WithError(err).Warn("create user request validation error")
		http.Error(w, "validation error", http.StatusBadRequest)
		return
	}

	newUser, err := h.userService.CreateNewUser(req)
	if err != nil {
		// для 500тых лучше Error т.к. сбой на стороне сервера
		logger.Log.WithError(err).Error("creating user problem")
		http.Error(w, "user creation error", http.StatusInternalServerError)
		return
	}

	// dto для ответа
	resp := dto.CreateUserResponse{
		ID:     newUser.ID,
		Name:   newUser.Name,
		Email:  newUser.Email,
		Role:   newUser.Role,
		Active: newUser.Active,
	}

	// Статус для создания — http.StatusCreated (201)
	w.Header().Set("Content-Type", "application/json")
	// добавляем Location header: по REST будет хорошим тоном для 201 Created
	w.Header().Set("Location", "/users/"+newUser.ID)
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logger.Log.WithError(err).Error("failed to encode create user response")
		return
	}

	// либо так
	//	result, err := json.Marshal(newUser)
	// if err != nil {
	// ...
	// что бы можно было изменить заголовки
	// если сериализовать не удасться
	// ...
	// } else { w.Write(result) }

}
