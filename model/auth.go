package model

import (
	"errors"

	"gorm.io/gorm"
)

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username" validate:"required,max=10"`
	Password string `json:"password" validate:"required,max=20"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	err := db.Where("username = ? AND password = ?", username, password).First(&auth).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}
