package dto

import (
	"errors"
	"strings"
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

type CreateRecordRequest struct {
	ID        string            `json:"id"`
	RecNumber uint              `json:"number"`
	RecType   models.RecordType `json:"type"`

	VisitorName      string `json:"visitorname"`
	VisitPurpose     string `json:"visitorporpuse"`
	OrganizationName string `json:"organizationname"`

	MalfunctionDesc       string
	MalfunctionCause      string
	MalfunctionResolution string
	MalfunctionStatus     string

	UserID   string
	UserName string

	StartTime time.Time
	EndTime   time.Time
}

func (crr *CreateRecordRequest) Normalize() {
	crr.ID = strings.ToLower(strings.TrimSpace(crr.ID))
	// TODO: add field normalization after the model is adjusted
}

func (crr *CreateRecordRequest) Validate() error {
	if crr.ID == "" {
		return errors.New("ID field could not be empty")
	}
	if crr.UserID == "" {
		return errors.New("User ID field could not be empty")
	}
	// TODO: add field validation after the model is adjusted.
	return nil
}

type CreateRecordResponse struct {
	ID        string            `json:"id"`
	RecNumber uint              `json:"number"`
	RecType   models.RecordType `json:"type"`

	VisitorName      string `json:"visitorname"`
	VisitPurpose     string `json:"visitorporpuse"`
	OrganizationName string `json:"organizationname"`

	MalfunctionDesc       string
	MalfunctionCause      string
	MalfunctionResolution string
	MalfunctionStatus     string

	UserID   string
	UserName string

	StartTime time.Time
	EndTime   time.Time
}
