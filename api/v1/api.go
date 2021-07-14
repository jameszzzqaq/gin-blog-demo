package v1

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func InitApi() {
	validate = validator.New()
}
