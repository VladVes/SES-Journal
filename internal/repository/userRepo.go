package repository

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/VladVes/SES-Journal/internal/models"
)

type UserRepository interface {
	GetUsersList() ([]models.User, error)
	CreateUser(user *models.User) (models.User, error)
	IsEmailExists(email string) (bool, error)
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(dbCon *gorm.DB) UserRepository {
	return &UserRepo{
		DB: dbCon,
	}
}

func (r *UserRepo) CreateUser(u *models.User) (models.User, error) {
	if err := r.DB.Create(u).Error; err != nil {
		return models.User{}, fmt.Errorf("user repo: %w", err)
	}
	return *u, nil
}

func (r *UserRepo) GetUsersList() ([]models.User, error) {
	var usersList []models.User

	if err := r.DB.Select("id", "login", "role").Find(&usersList).Error; err != nil {
		return nil, err
	}

	return usersList, nil
}

func (r *UserRepo) IsEmailExists(email string) (bool, error) {
	// TODO: implement real logic

	// simplified result data
	return email == "ivan@mail.ru", nil

}
