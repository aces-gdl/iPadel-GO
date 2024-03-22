package models

import (
	"gorm.io/gorm"
)

type TournamentGroup struct {
	gorm.Model
	Name         string
	GroupNumber  int
	CategoryID   uint
	TournamentID uint
}
