package models

import (
	"gorm.io/gorm"
)

type TournamentCategories struct {
	gorm.Model
	TournamentID uint
	Tournament   Tournament `gorm:"foreignKey:TournamentID"`
	CategoryID   uint
	Category     Category `gorm:"foreignKey:CategoryID"`
	Active       bool
}
