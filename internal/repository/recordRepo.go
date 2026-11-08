package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/VladVes/SES-Journal/internal/dto"
	"github.com/VladVes/SES-Journal/internal/models"
)

type RecordRepository interface {
	GetRecordsList() ([]dto.RecordListItem, error)
	GetRecordByID(id string) (*models.Record, error)
}

type RecordRepo struct {
	DB *gorm.DB
}

func NewRecordRepo(dbCon *gorm.DB) *RecordRepo {
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

func (rr *RecordRepo) GetRecordByID(id string) (*models.Record, error) {
	var record models.Record

	if err := rr.DB.Preload("User").Where("ID = ?", id).First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFoundByID
		}
		return nil, err
	}

	return &record, nil
}
