package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model   `gorm:"embedded"`
	Email        string `gorm:"unique"`
	Password     string
	GoogleID     string
	ImageURL     string
	HasPicture   int
	Name         string
	FamilyName   string
	GivenName    string
	PermissionID uint
	Ranking      int
	CategoryID   uint
	Birthday     time.Time
	Phone        string
	MemberSince  time.Time
	Active       bool
}
