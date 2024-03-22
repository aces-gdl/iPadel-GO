package models

import (
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	Description string
	Active      bool
}
