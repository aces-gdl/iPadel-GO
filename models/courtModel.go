package models

import (
	"gorm.io/gorm"
)

type Court struct {
	gorm.Model
	Name    string
	Indoors bool
	ClubID  uint
}
