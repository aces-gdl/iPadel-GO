package models

import (
	"gorm.io/gorm"
)

type TournamentTeamByGroup struct {
	gorm.Model
	Name         string
	GroupNumber  int
	TournamentID uint
	CategoryID   uint
	GroupID      uint
	TeamID       uint
}
