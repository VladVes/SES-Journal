package services

import (
	"github.com/VladVes/SES-Journal/internal/dto"
	"github.com/VladVes/SES-Journal/internal/models"
	"github.com/VladVes/SES-Journal/internal/repository"
)

type RecordService struct {
	repo repository.RecordRepository
}

func NewRecordSerivce(recordRepo repository.RecordRepository) *RecordService {
	return &RecordService{
		repo: recordRepo,
	}
}

func (rs *RecordService) GetRecordsList() ([]dto.RecordListItem, error) {
	return rs.repo.GetRecordsList()
}

func (rs *RecordService) GetRecordByID(id string) (*models.Record, error) {
	return rs.repo.GetRecordByID(id)
}
