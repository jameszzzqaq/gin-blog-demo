package service

import (
	"errors"

	"github.com/yu1er/gin-blog/model"
	"github.com/yu1er/gin-blog/model/request"
	"gorm.io/gorm"
)

func GetArticlesPage(info request.ArticleListGet) ([]model.Article, int, error) {
	var total int64
	var arts []model.Article

	limit := info.PageSize
	offset := info.PageSize * (info.PageNum - 1)

	sql := db.Model(&model.Article{}).Preload("Tag")
	if info.Title != "" {
		sql = sql.Where("title like ?", "%"+info.Title+"%")
	}

	if info.Desc != "" {
		sql = sql.Where("desc like ?", "%"+info.Desc+"%")
	}

	if info.Content != "" {
		sql = sql.Where("content like %", "%"+info.Content+"%")
	}

	if info.TagID != nil {
		sql = sql.Where("tag_id = ?", info.TagID)
	}

	_ = sql.Count(&total).Error

	err := sql.Offset(offset).Limit(limit).Find(&arts).Error
	return arts, int(total), err
}

func GetArticlesCount() int {
	var count int64
	db.Model(&model.Article{}).Count(&count)
	return int(count)
}

// TODO: bug, need to get tag meanwhile.
func GetArticleById(id int) (a model.Article, err error) {
	err = db.Preload("Tag").First(&a, id).Error
	return
}

func CheckArticlExistById(id int) bool {
	err := db.First(&model.Article{}, id).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func AddArticle(req *request.AriticleAdd) error {
	article := model.Article{
		TagID:      req.TagID,
		Title:      req.Title,
		Desc:       req.Desc,
		Content:    req.Content,
		CreatedBy:  req.CreatedBy,
		ModifiedBy: req.CreatedBy,
		State:      *req.State,
	}
	err := db.Create(article).Error
	return err
}

func UpdateArticle(id int, req *request.ArticleUpdate) error {
	err := db.First(&model.Article{}, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	artMap := map[string]interface{}{
		"TagID":      req.TagID,
		"Title":      req.Title,
		"Desc":       req.Desc,
		"Content":    req.Content,
		"ModifiedBy": req.ModifiedBy,
		"State":      req.State,
	}

	err = db.Model(&model.Article{}).Updates(&artMap).Error
	return err
}

func DeleteArticle(id int) error {
	err := db.First(&model.Article{}, id).Error
	if err != nil {
		return err
	}

	err = db.Delete(&model.Article{}, id).Error
	return err
}
