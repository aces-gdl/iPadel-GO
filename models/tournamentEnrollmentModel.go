package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TournamentEnrollment struct {
	ID           uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	UserID       uuid.UUID
	User         User       `gorm:"foreignKey:UserID;references:ID"`
	TournamentID uuid.UUID  `gorm:"type:uuid;"`
	Tournament   Tournament `gorm:"foreignKey:TournamentID;references:ID"`
	CategoryID   uuid.UUID  `gorm:"type:uuid;"`
	Category     Category   `gorm:"foreignKey:CategoryID;references:ID"`
	Active       bool
}
