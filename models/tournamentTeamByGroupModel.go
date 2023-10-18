package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TournamentTeamByGroup struct {
	ID           uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Name         string
	GroupNumber  int
	TournamentID uuid.UUID `gorm:"type:uuid;"`
	CategoryID   uuid.UUID `gorm:"type:uuid;"`
	GroupID      uuid.UUID `gorm:"type:uuid;"`
	TeamID       uuid.UUID `gorm:"type:uuid;"`
}
