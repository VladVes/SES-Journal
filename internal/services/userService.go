package services

import (
	"fmt"
	"time"

	"github.com/VladVes/SES-Journal/internal/models"
	"github.com/VladVes/SES-Journal/internal/repository"

	"github.com/VladVes/SES-Journal/internal/dto"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		repo: userRepo,
	}
}

func (s *UserService) GetUsersList() ([]models.User, error) {
	return s.repo.GetUsersList()
}

func (s *UserService) CreateNewUser(in dto.CreateUserRequest) (models.User, error) {

	// ID       string `gorm:"primaryKey"`
	// Name     string `gorm:"column:login;size:30;not null"`
	// Email    string `gorm:"type:varchar(100)"`
	// Password string `gorm:"not nul"`
	// Role     string
	// Active   bool

	// Records []Record
	// // Comments []Comment

	// LastLogin time.Time
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt time.Time

	// Пример бизнес-проверки
	// Можно сначала проверить через репозиторий, есть ли юзер с таким email
	// техническую ошибку (упал коннект к БД) мы просто оборачиваем через %w,
	// чтобы в логах был стек/контекст.
	// Бизнес-ошибку мы возвращаем как есть в виде ErrEmailExists,
	// чтобы хендлер мог по ней понять, что это именно конфликт.
	exists, err := s.repo.IsEmailExists(in.Email)
	if err != nil {
		// Если это техническая ошибка репозитория - оборачиваем, но оставляем контекст
		// т.е. ошибка при работе с базой при попытке проверить email
		return models.User{}, fmt.Errorf("check email exists: %w", err)
	}
	if exists {
		// Бизнес-ошибка - возвращаем именно наш sentinel
		return models.User{}, ErrEmailExists
	}


	tmp := models.User{
		ID:       "2342",
		Name:     "Ivan",
		Email:    "ivan@mail.ru",
		Password: "sfdsdf334",
		Role:     "egineer",
		Active:   true,
		Records:  []models.Record{},

		LastLogin: time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Now(),
	}
	return tmp, nil
}
