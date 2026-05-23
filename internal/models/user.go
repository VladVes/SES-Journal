package models

import (
	"time"
)

type User struct {
	ID       string `gorm:"primaryKey"`
	Name     string `gorm:"column:login;size:30;not null"`
	Email    string `gorm:"type:varchar(100)"`
	Password string `gorm:"not nul"`
	Role     string
	Active   bool

	Records []LogRecord
	// Comments []Comment

	LastLogin time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
