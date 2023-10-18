package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TournamentTeam struct {
	ID           uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Name         string         `gorm:"unique"`
	CategoryID   uuid.UUID      `gorm:"type:uuid;"`
	Member1ID    uuid.UUID      `gorm:"type:uuid;"`
	Name1        string
	Ranking1     int
	Member2ID    uuid.UUID `gorm:"type:uuid;"`
	Name2        string
	Ranking2     int
	TournamentID uuid.UUID `gorm:"type:uuid;"`
}
