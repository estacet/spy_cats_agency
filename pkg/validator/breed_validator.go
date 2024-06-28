package validator

import (
	"github.com/go-playground/validator/v10"
	"spy-cats/pkg/catapi"
)

type BreedValidator struct {
	catApiClient *catapi.Client
}

func NewBreedValidator(catApiClient *catapi.Client) *BreedValidator {
	return &BreedValidator{catApiClient: catApiClient}
}

func (v *BreedValidator) Validate(fl validator.FieldLevel) bool {
	breeds, err := v.catApiClient.GetBreeds()
	if err != nil {
		// TODO: log error
		return false
	}

	for _, breed := range breeds {
		if breed.Id == fl.Field().String() {
			return true
		}
	}

	return false
}
