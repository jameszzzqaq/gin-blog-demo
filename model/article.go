package model

import (
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
