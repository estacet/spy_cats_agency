package model

import (
	"github.com/google/uuid"
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

	return c, nil
}

func (c *SpyCat) Update(salary float64) {
	c.Salary = salary
}
