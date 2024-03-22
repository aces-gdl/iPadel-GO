package models

import (
	"time"

	"gorm.io/gorm"
)

type TournamentTimeSlots struct {
	gorm.Model
	TournamentID uint
	CategoryID   uint
	Description  string
	CourtNumber  int
	StartTime    time.Time
	EndTime      time.Time
	GameID       uint
	Taken        bool
	Active       bool
}
