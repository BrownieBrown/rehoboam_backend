package model

import (
	"gorm.io/gorm"
	"time"
)

type Entry struct {
	gorm.Model
	CreatedAt time.Time
	UpdatedAt time.Time
	UserId    uint
}
