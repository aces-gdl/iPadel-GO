package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TournamentGroup struct {
	ID           uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Name         string         `gorm:"unique"`
	GroupNumber  int
	CategoryID   uuid.UUID `gorm:"type:uuid;"`
	TournamentID uuid.UUID `gorm:"type:uuid;"`
}
