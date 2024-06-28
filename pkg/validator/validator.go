package validator

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

func New(breedValidator *BreedValidator) (*validator.Validate, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("json")
	})

	err := validate.RegisterValidation("breedValidator", breedValidator.Validate)
	if err != nil {
		return nil, err
	}

	return validate, nil
}
