package model

import "gorm.io/gorm"

type SaveData struct {
	*gorm.Model
	UserID   uint   `gorm:"not null"`
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
	Platform string `gorm:"not null"`
}

type DataUser struct {
	*gorm.Model
	Username string
	Password string
}
