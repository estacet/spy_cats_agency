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

func NewSpyCat(name string, yearsOfExperience int, breed string, salary float64) (*SpyCat, error) {
	c := &SpyCat{
		ID:                uuid.New(),
		Name:              name,
		YearsOfExperience: yearsOfExperience,
		Breed:             breed,
		Salary:            salary,
	}

	err := c.validateBreed()

	if err != nil {
		fmt.Errorf("cat %v was not registered. The reason: %v", c.Name, err)
		return nil, err
	}

	return c, nil
}

func (c *SpyCat) validateBreed() error {
	breeds, err := GetBreeds()

	if err != nil {
		return errors.New("unable to fetch breeds info")
	}

	for _, breed := range *breeds {
		if breed.Id == c.Breed {
			fmt.Printf("Breed of agent %c approved", c.Name)
			return nil
		}
		err = apperror.NewBreedNotFoundError("The breed " + c.Breed + " is not registered")
	}

	return err
}

func (c *SpyCat) Update(salary float64) {
	c.Salary = salary
}
