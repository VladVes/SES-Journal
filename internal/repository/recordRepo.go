package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/VladVes/SES-Journal/internal/dto"
	"github.com/VladVes/SES-Journal/internal/models"
)

type RecordRepository interface {
	GetRecordsList() ([]dto.RecordListItem, error)
	GetRecordByID(id string) (models.Record, error)
	CreateNewRecord(rec *models.Record) (models.Record, error)
	IsUserExists(userID string) (bool, error)
}

type RecordRepo struct {
	DB *gorm.DB
}

func NewRecordRepo(dbCon *gorm.DB) RecordRepository {
	return &RecordRepo{
		DB: dbCon,
	}
}

func (rr *RecordRepo) GetRecordsList() ([]dto.RecordListItem, error) {
	var recordsList []dto.RecordListItem

	if err := rr.DB.Model(models.Record{}).Select("ID", "RecNumber", "RecType", "StartTime", "EndTime").Scan(&recordsList).Error; err != nil {
		return nil, err
	}

	return recordsList, nil
}

func (rr *RecordRepo) GetRecordByID(id string) (models.Record, error) {
	var record models.Record

	if err := rr.DB.Preload("User").Where("ID = ?", id).First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Record{}, ErrRecordNotFoundByID
		}
		return models.Record{}, err
	}

	return record, nil
}

func (rr *RecordRepo) CreateNewRecord(rec *models.Record) (models.Record, error) {
	if err := rr.DB.Create(rec).Error; err != nil {
		return models.Record{}, fmt.Errorf("records repo: %w", err)
	}
	return *rec, nil
}

func (rr *RecordRepo) IsUserExists(userID string) (bool, error) {
	if err := rr.DB.Select("id").First(&models.User{}, "id = ?", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
