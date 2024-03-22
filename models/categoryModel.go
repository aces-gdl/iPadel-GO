package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Description string
	Color       string
	Level       int
	Active      bool
}
