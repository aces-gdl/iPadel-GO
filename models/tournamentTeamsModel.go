package models

import (
	"gorm.io/gorm"
)

type TournamentTeam struct {
	gorm.Model
	Name         string
	CategoryID   uint
	Member1ID    uint
	Name1        string
	Ranking1     int
	Member2ID    uint
	Name2        string
	Ranking2     int
	TournamentID uint
}
