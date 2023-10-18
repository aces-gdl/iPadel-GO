package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TournamentGames struct {
	ID                    uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt `gorm:"index"`
	TournamentID          uuid.UUID      `gorm:"type:uuid;"`
	CategoryID            uuid.UUID      `gorm:"type:uuid;"`
	TournamentGroupID     uuid.UUID      `gorm:"type:uuid;"`
	Team1ID               uuid.UUID      `gorm:"type:uuid;"`
	Team2ID               uuid.UUID      `gorm:"type:uuid;"`
	TournamentTimeSlotsID uuid.UUID      `gorm:"type:uuid;"`
	Active                bool
}
