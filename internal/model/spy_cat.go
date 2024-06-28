package model

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"spy-cats/pkg/apperror"
)

type SpyCat struct {
	ID                uuid.UUID `json:"id"`
	Name              string    `json:"name"`
	YearsOfExperience int       `json:"years_of_experience"`
	Breed             string    `json:"breed"`
	Salary            float64   `json:"salary"`
}

func NewSpyCat(id uuid.UUID, name string, yearsOfExperience int, breed string, salary float64) *SpyCat {
	c := &SpyCat{
		ID:                id,
		Name:              name,
		YearsOfExperience: yearsOfExperience,
		Salary:            salary,
		Breed:             breed,
	}

	err := c.validateBreed()
	if err != nil {
		fmt.Errorf("cat %v was not registered. The reason: %v", c.Name, err)
	}

	return c
}

func (c *SpyCat) validateBreed() error {
	breeds, err := GetBreeds()

	if err != nil {
		return errors.New("unable to fetch breeds info")
	}

	for _, breed := range *breeds {

		if breed.id == c.Breed {
			fmt.Printf("Breed of agent %c approved", c.Name)
		}

		return apperror.NewBreedNotFoundError("The breed " + c.Breed + " is not registered")
	}
	return nil
}
