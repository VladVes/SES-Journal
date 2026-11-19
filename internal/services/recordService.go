package services

import (
	"fmt"

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

func (rs *RecordService) GetRecordByID(id string) (models.Record, error) {
	return rs.repo.GetRecordByID(id)
}

func (rs *RecordService) CreateNewRecord(in *dto.CreateRecordRequest) (*models.Record, error) {

	// TODO: should we check record id for existens by the repo methods?

	exists, err := rs.repo.IsUserExists(in.UserID)
	if err != nil {
		return nil, fmt.Errorf("while trying to check is user exists: %w\n", err)
	}
	if !exists {
		return nil, ErrUserNotFound
	}

	newRecord := models.Record{
		ID:        in.ID,
		RecNumber: in.RecNumber,
		RecType:   in.RecType,

		MalfunctionDesc:       in.MalfunctionDesc,
		MalfunctionResolution: in.MalfunctionResolution,
		MalfunctionCause:      in.MalfunctionCause,
		MalfunctionStatus:     in.MalfunctionStatus,

		UserID: in.UserID,

		VisitorName:      in.VisitorName,
		VisitPurpose:     in.VisitPurpose,
		OrganizationName: in.OrganizationName,

		StartTime: in.StartTime,
		EndTime:   in.EndTime,
	}

	record, err := rs.repo.CreateNewRecord(&newRecord)
	if err != nil {
		return &record, err
	}

	return &newRecord, nil
}
