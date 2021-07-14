package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	Model

	Name       string `json:"name" validate:"max=100"`
	State      int    `json:"state" validate:"oneof=0 1"`
	CreatedBy  string `json:"created_by" validate:"max=100"`
	ModifiedBy string `json:"modified_by" validate:"max=100"`
}

func (t *Tag) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("CreatedOn", time.Now().Unix())
	tx.Statement.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func (t *Tag) BeforeUpdate(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func GetTagsPage(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagById(id int) (tag Tag) {
	db.First(&tag, id)
	return
}

func GetTagsCount(maps interface{}) int {
	var count int64
	db.Model(&Tag{}).Where(maps).Count(&count)

	return int(count)
}

func CheckTagExistByName(name string) bool {
	err := db.Where("name = ?", name).Take(&Tag{}).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func CheckTagExistById(id int) bool {
	err := db.First(&Tag{}, id).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func AddTag(name string, state int, createdBy string) {
	db.Create(&Tag{
		Name:       name,
		State:      state,
		CreatedBy:  createdBy,
		ModifiedBy: createdBy,
	})
}

func UpdateTag(id int, tag Tag) bool {
	db.Model(&tag).Updates(tag)
	return true
}

func DeleteTag(id int) bool {
	db.Delete(&Tag{}, id)
	return true
}
