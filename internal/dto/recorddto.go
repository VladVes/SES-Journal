package dto

import (
	"time"

	"github.com/VladVes/SES-Journal/internal/models"
)

type RecordListItem struct {
	ID        string
	RecNumber uint
	RecType   models.RecordType
	StartTime time.Time
	EndTime   time.Time
}

type GetRecordResponse struct {
	ID        string
	RecNumber uint
	RecType   models.RecordType

	VisitorName      string
	VisitPurpose     string
	OrganizationName string

	MalfunctionDesc       string
	MalfunctionCause      string
	MalfunctionResolution string
	MalfunctionStatus     string

	UserID   string
	UserName string

	StartTime time.Time
	EndTime   time.Time
}
