package services

import (
	"github.com/VladVes/SES-Journal/internal/logger"
	"github.com/VladVes/SES-Journal/internal/models"
	"github.com/VladVes/SES-Journal/internal/repository"
)

type UserService struct {
	repo *repository.UsersRepo
}

func NewUserServise() *UserService {
	userRepo := repository.NewUsersRepo()
	return &UserService{
		repo: userRepo,
	}
}

func (s *UserService) GetUsersList() ([]models.User, error) {
	users, err := s.repo.GetUsersList()
	if err != nil {
		logger.Log.WithError(err).Fatal("error while getting users list")
	}

	return users, nil
}
