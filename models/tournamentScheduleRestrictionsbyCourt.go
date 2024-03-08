package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScheduleRestrictionsByTeam struct {
	ID               uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	TournamentID     uuid.UUID
	BlockedStartTime time.Time
	BlockedEndTime   time.Time
	CapturedBy       uuid.UUID
	Type             int //1 by Team, 2 by Court
	Active           bool
}
