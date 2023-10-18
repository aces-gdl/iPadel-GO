package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tournament struct {
	ID               uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	Description      string
	StartDate        time.Time
	EndDate          time.Time
	StartTime        time.Time
	EndTime          time.Time
	RoundrobinDays   int
	PlayOffDays      int
	HostClubID       uuid.UUID `gorm:"type:uuid;"`
	Club             Club      `gorm:"foreignKey:HostClubID;references:ID"`
	GameDuration     int
	RoundrobinCourts int
	PlayoffCourts    int
	Active           bool
}
