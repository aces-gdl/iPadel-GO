package models

import (
	"time"

	"gorm.io/gorm"
)

type ScheduleRestrictionsByTeam struct {
	gorm.Model
	TournamentID     uint
	BlockedStartTime time.Time
	BlockedEndTime   time.Time
	CapturedBy       uint
	Type             int //1 by Team, 2 by Court
	Active           bool
}
