package models

import (
	"time"

	"gorm.io/gorm"
)

type Tournament struct {
	gorm.Model
	Description      string
	StartDate        time.Time
	EndDate          time.Time
	StartTime        time.Time
	EndTime          time.Time
	RoundrobinDays   int
	PlayOffDays      int
	HostClubID       uint
	Club             Club `gorm:"foreignKey:HostClubID;references:ID"`
	GameDuration     int
	RoundrobinCourts int
	PlayoffCourts    int
	Active           bool
}
