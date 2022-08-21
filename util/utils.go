package util

import "github.com/go-playground/validator/v10"

type InitValidator struct {
	*validator.Validate
}

func NewInitValidator() *validator.Validate {
	v := validator.New()
	return v
}
