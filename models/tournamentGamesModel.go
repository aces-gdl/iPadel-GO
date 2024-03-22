package models

import (
	"gorm.io/gorm"
)

type TournamentGames struct {
	gorm.Model
	TournamentID          uint
	CategoryID            uint
	TournamentGroupID     uint
	Team1ID               uint
	Team2ID               uint
	TournamentTimeSlotsID uint
	GameType              string
	Comment               string
	Active                bool
}
