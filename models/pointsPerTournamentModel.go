package models

import (
	"gorm.io/gorm"
)

type PointsPerTournament struct {
	gorm.Model

	Active bool
}
