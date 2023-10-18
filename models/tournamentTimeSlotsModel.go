package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TournamentTimeSlots struct {
	ID           uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	TournamentID uuid.UUID      `gorm:"type:uuid;"`
	CategoryID   uuid.UUID      `gorm:"type:uuid;"`
	Description  string
	CourtNumber  int
	StartTime    time.Time
	EndTime      time.Time
	GameID       uuid.UUID `gorm:"type:uuid;"`
	Taken        bool
	Active       bool
}
