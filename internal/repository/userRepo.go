package repository

import (
	"gorm.io/gorm"

	"github.com/VladVes/SES-Journal/internal/models"
)

type UserRepository interface {
	GetUsersList() ([]models.User, error)
	CreateUser(user *models.User) (string, error)
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(dbCon *gorm.DB) UserRepository {
	return &UserRepo{
		DB: dbCon,
	}
}

func (r *UserRepo) CreateUser(u *models.User) (string, error) {
	// TODO
	return "", nil
}

func (r *UserRepo) GetUsersList() ([]models.User, error) {
	var usersList []models.User

	if err := r.DB.Select("id", "login", "role").Find(&usersList).Error; err != nil {
		return nil, err
	}

	return usersList, nil
}
