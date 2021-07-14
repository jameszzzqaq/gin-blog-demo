package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// gorm hooks
func (a *Article) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("CreatedOn", time.Now().Unix())
	tx.Statement.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func (a *Article) BeforeUpdate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func GetArticlesPage(pageNum int, pageSize int) (as []Article) {
	db.Preload("Tag").Offset(pageNum).Limit(pageSize).Find(&as)
	return
}

func GetArticlesCount() int {
	var count int64
	db.Model(&Article{}).Count(&count)
	return int(count)
}

func GetArticleById(id int) (a Article) {
	db.First(&a)
	return
}

func CheckArticlExistById(id int) bool {
	err := db.First(&Article{}, id).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func AddArticle(article Article) bool {
	db.Create(&article)
	return true
}

func UpdateArticle(id int, article Article) bool {
	article.ID = id
	db.Model(&article).Updates(&article)
	return true
}

func DeleteArticle(id int) bool {
	db.Delete(&Article{}, id)
	return true
}
