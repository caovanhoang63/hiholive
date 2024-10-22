package core

import (
	"github.com/go-playground/validator/v10"
)

var Validator = &validate{
	validator.New(),
}

type validate struct {
	validator *validator.Validate
}

func (v *validate) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func (v *validate) ValidateField(i interface{}) (string, error) {
	err := v.validator.Struct(i)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return err.StructField(), err
		}
	}

	return "", nil
}
