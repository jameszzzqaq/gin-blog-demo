package service

import (
	"errors"

	"github.com/yu1er/gin-blog/model"
	"gorm.io/gorm"
)

func GetArticlesPage(pageNum int, pageSize int) (as []model.Article) {
	db.Preload("Tag").Offset(pageNum).Limit(pageSize).Find(&as)
	return
}

func GetArticlesCount() int {
	var count int64
	db.Model(&model.Article{}).Count(&count)
	return int(count)
}

// TODO: bug, need to get tag meanwhile.
func GetArticleById(id int) (a model.Article) {
	db.First(&a)
	return
}

func CheckArticlExistById(id int) bool {
	err := db.First(&model.Article{}, id).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func AddArticle(article model.Article) bool {
	db.Create(&article)
	return true
}

func UpdateArticle(id int, article model.Article) bool {
	article.ID = id
	db.Model(&article).Updates(&article)
	return true
}

func DeleteArticle(id int) bool {
	db.Delete(&model.Article{}, id)
	return true
}
