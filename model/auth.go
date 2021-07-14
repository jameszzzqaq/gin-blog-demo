package model

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username" validate:"required,max=10"`
	Password string `json:"password" validate:"required,max=20"`
}
