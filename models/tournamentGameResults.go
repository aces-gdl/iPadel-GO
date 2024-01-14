package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TournamentGameResult struct {
	ID            uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	GameID        uuid.UUID      `gorm:"type:uuid;"`
	Team1Set1     int
	Team1Set2     int
	Team1Set3     int
	Team2Set1     int
	Team2Set2     int
	Team2Set3     int
	Winner        int
	WinningReason string
	Comments      string
}
