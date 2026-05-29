package repository

import (
	"gorm.io/gorm"

	"github.com/VladVes/SES-Journal/internal/models"
)

type UserRepository interface {
	GetUsersList() (models.User, error)
	CreateUser(user *models.User) (string, error)
}

type UsersRepo struct {
	DB *gorm.DB
}

func NewUsersRepo() *UsersRepo {
	return &UsersRepo{
		DB: GetDBConnection(),
	}
}

func (r *UsersRepo) CreateUser(u *models.User) (string, error) {
	// TODO
	return "", nil
}

func (r *UsersRepo) GetUsersList() ([]models.User, error) {
	var usersList []models.User

	if err := r.DB.Select("id", "login", "role").Find(&usersList).Error; err != nil {
		return nil, err
	}

	return usersList, nil
}
