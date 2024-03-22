package models

import (
	"gorm.io/gorm"
)

type Club struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Description string
	Contact     string
	ImageURL    string
	Address     string
	Phone       string
	Courts      []Court `gorm:"foreignKey:ClubID"`
}
