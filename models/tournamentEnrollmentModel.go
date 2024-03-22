package models

import (
	"gorm.io/gorm"
)

type TournamentEnrollment struct {
	gorm.Model
	UserID       uint
	User         User `gorm:"foreignKey:UserID;references:ID"`
	TournamentID uint
	Tournament   Tournament `gorm:"foreignKey:TournamentID;references:ID"`
	CategoryID   uint
	Category     Category `gorm:"foreignKey:CategoryID;references:ID"`
	Active       bool
}
