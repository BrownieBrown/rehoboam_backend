package models

import (
	"rehoboam/pkg/database"

	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	UserId uint
}

func (entry *Entry) Save() (*Entry, error) {
	err := database.Connect().Create(&entry).Error
	if err != nil {
		return &Entry{}, err
	}
	return entry, nil
}
