package service

import (
	"errors"

	"github.com/yu1er/gin-blog/model"
	"gorm.io/gorm"
)

func CheckAuthExist(a *model.Auth) bool {
	var auth = &model.Auth{}
	err := db.First(auth, "username = ? AND password = ?", a.Username, a.Password).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}
