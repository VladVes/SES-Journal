package repository

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/VladVes/SES-Journal/internal/dto"
	"github.com/VladVes/SES-Journal/internal/models"
)

type RecordRepository interface {
	GetRecordsList() ([]dto.RecordListItem, error)
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
	var recordsList []models.Record
	recordListItem := []dto.RecordListItem{
		{ID: "wfs", RecNumber: 333, RecType: "sdfsfd", StartTime: time.Now(), EndTime: time.Now()},
	}

	// if err := rr.DB.Select("rec_number", "rec_type", "start_time", "end_time").Find(&recordsList).Error; err != nil {
	// 	return nil, err
	// }
	// убрать кучу полей с датами
	if err := rr.DB.Select("rec_number", "rec_type").Find(&recordsList).Error; err != nil {
		return nil, err
	}
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println(recordsList)

	return recordListItem, nil

}
