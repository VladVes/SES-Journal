package models

import "time"

type LogType string

const (
	Incident      LogType = "incident"
	ShiftWork     LogType = "shiftWork"
	ShiftInfo     LogType = "shiftInfo"
	VisitRegisty  LogType = "visitRegistry"
	CustomEntries LogType = "customEntries"
)

type LogRecord struct {
	ID        string `gorm:"primaryKey"`
	RecNumber uint
	RecType   LogType

	VisitorName     string
	VisitPoupose    string
	OranizationName string

	MalfunctionDesc       string
	MalfunctionCause      string
	MalfunctionResolution string
	MalfunctionStatus     string

	UserID string
	User   User `gorm:"foreignKey:UserID"`

	StartTime time.Time
	EndTime   time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
