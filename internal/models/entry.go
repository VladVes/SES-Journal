package models

import "time"

type EntryType string

const (
	Incident      EntryType = "incident"
	ShiftWork     EntryType = "shiftWork"
	ShiftInfo     EntryType = "shiftInfo"
	VisitRegisty  EntryType = "visitRegistry"
	CustomEntries EntryType = "customEntries"
)

type Entry struct {
	ID        string `gorm:"primaryKey"`
	RecNumber uint
	RecType   EntryType

	VisitorName      string
	VisitPoupose     string
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
