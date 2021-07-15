package service

import (
	"errors"

	"github.com/yu1er/gin-blog/model"
	"github.com/yu1er/gin-blog/model/request"
	"gorm.io/gorm"
)

func GetTagsPage(info request.TagListGet) ([]model.Tag, int, error) {
	var total int64
	var tags []model.Tag

	limit := info.PageSize
	offset := info.PageSize * (info.PageNum - 1)

	sql := db.Model(&model.Tag{})
	if info.Name != "" {
		sql.Where("name Like ?", "%"+info.Name+"%")
	}

	_ = sql.Count(&total).Error

	err := sql.Offset(offset).Limit(limit).Find(&tags).Error
	return tags, int(total), err
}

func GetTagById(id int) (tag model.Tag, err error) {
	err = db.First(&tag, id).Error
	return tag, err
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

func AddTag(tp *model.Tag) error {
	err := db.Create(tp).Error
	return err
}

func UpdateTag(id int, tag *model.Tag) error {
	err := db.Model(&model.Tag{}).Where("id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	tagMap := map[string]interface{}{
		"Name":       tag.Name,
		"State":      tag.State,
		"ModifiedBy": tag.ModifiedBy,
	}

	err = db.Model(&model.Tag{}).Where("id = ?", id).Updates(tagMap).Error
	return err
}

func DeleteTag(id int) error {
	err := db.Delete(&model.Tag{}, id).Error
	return err
}
