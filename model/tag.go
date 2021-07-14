package model

import (
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
