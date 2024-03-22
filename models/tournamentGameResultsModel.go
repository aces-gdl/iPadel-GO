package models

import (
	"gorm.io/gorm"
)

type TournamentGameResult struct {
	gorm.Model
	GameID        uint
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
