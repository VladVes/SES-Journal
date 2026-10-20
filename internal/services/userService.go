package services

import (
	"github.com/VladVes/SES-Journal/internal/models"
	"github.com/VladVes/SES-Journal/internal/repository"
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
