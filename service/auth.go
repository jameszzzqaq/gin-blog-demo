package service

import (
	"errors"

	"github.com/yu1er/gin-blog/model"
	"gorm.io/gorm"
)

func CheckAuth(username, password string) bool {
	var auth model.Auth
	err := db.Where("username = ? AND password = ?", username, password).First(&auth).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}
