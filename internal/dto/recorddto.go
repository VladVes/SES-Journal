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
