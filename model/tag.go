package model

import (
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	Model

	Name       string `json:"name"`
	State      int    `json:"state"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
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
