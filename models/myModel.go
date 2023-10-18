package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MyModel struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
