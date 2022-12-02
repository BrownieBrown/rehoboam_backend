package model

import (
	"gorm.io/gorm"
	"rehoboam/database"
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
