package models

import "time"

type RecordType string

const (
	Incident      RecordType = "incident"
	ShiftWork     RecordType = "shiftWork"
	ShiftInfo     RecordType = "shiftInfo"
	VisitRegisty  RecordType = "visitRegistry"
	CustomEntries RecordType = "customEntries"
)

type Record struct {
	ID        string `gorm:"primaryKey"`
	RecNumber uint
	RecType   RecordType

	VisitorName      string
	VisitPurpose     string
	OrganizationName string

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
