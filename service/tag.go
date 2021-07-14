package service

import (
	"errors"

	"github.com/yu1er/gin-blog/model"
	"gorm.io/gorm"
)

func GetTagsPage(pageNum int, pageSize int, maps interface{}) (tags []model.Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagById(id int) (tag model.Tag) {
	db.First(&tag, id)
	return
}

func GetTagsCount(maps interface{}) int {
	var count int64
	db.Model(&model.Tag{}).Where(maps).Count(&count)

	return int(count)
}

func CheckTagExistByName(name string) bool {
	err := db.Where("name = ?", name).Take(&model.Tag{}).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func CheckTagExistById(id int) bool {
	err := db.First(&model.Tag{}, id).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func AddTag(name string, state int, createdBy string) {
	db.Create(&model.Tag{
		Name:       name,
		State:      state,
		CreatedBy:  createdBy,
		ModifiedBy: createdBy,
	})
}

func UpdateTag(id int, tag model.Tag) bool {
	db.Model(&tag).Updates(tag)
	return true
}

func DeleteTag(id int) bool {
	db.Delete(&model.Tag{}, id)
	return true
}
