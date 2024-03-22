package models

import (
	"gorm.io/gorm"
)

type PlayersByTournament struct {
	gorm.Model    `gorm:"embedded"`
	TournamentID  uint
	CategoryID    uint
	UserID        uint
	PaymentStatus string
}
