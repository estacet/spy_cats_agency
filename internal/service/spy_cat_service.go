package service

import (
	"context"
	"github.com/google/uuid"
	"spy-cats/internal/model"
	"spy-cats/internal/repository"
)

type CreateSpyCatArgs struct {
	Name              string  `json:"name"`
	YearsOfExperience int     `json:"years_of_experience"`
	Breed             string  `json:"breed" validate:"breedValidator"`
	Salary            float64 `json:"salary"`
}

type UpdateSpyCatArgs struct {
	Salary float64 `json:"salary"`
}

type SpyCatDetails struct {
	ID                uuid.UUID `json:"id"`
	Name              string    `json:"name"`
	YearsOfExperience int       `json:"years_of_experience"`
	Breed             string    `json:"breed"`
	Salary            float64   `json:"salary"`
	//Missions          *[]model.Mission `json:"missions"`
}

type SpyCatService struct {
	spyCatRepository *repository.SpyCatRepository
}

func NewSpyCatService(spyCatRepository *repository.SpyCatRepository) *SpyCatService {
	return &SpyCatService{spyCatRepository: spyCatRepository}
}

func (s *SpyCatService) Create(ctx context.Context, args *CreateSpyCatArgs) error {
	spyCat, err := model.NewSpyCat(
		args.Name,
		args.YearsOfExperience,
		args.Breed,
		args.Salary,
	)
	if err != nil {
		return err
	}

	return s.spyCatRepository.Create(ctx, spyCat)
}

func (s *SpyCatService) GetById(ctx context.Context, id uuid.UUID) (*SpyCatDetails, error) {
	spyCat, err := s.spyCatRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	requestedSpyCat := &SpyCatDetails{
		spyCat.ID,
		spyCat.Name,
		spyCat.YearsOfExperience,
		spyCat.Breed,
		spyCat.Salary,
	}

	return requestedSpyCat, nil
}

func (s *SpyCatService) GetList(ctx context.Context) ([]*SpyCatDetails, error) {
	spyCats, err := s.spyCatRepository.GetList(ctx)
	if err != nil {
		return nil, err
	}

	spyCatsList := make([]*SpyCatDetails, len(spyCats))

	for i, spyCat := range spyCats {
		cat := &SpyCatDetails{
			ID:                spyCat.ID,
			Name:              spyCat.Name,
			YearsOfExperience: spyCat.YearsOfExperience,
			Breed:             spyCat.Breed,
			Salary:            spyCat.Salary,
		}
		spyCatsList[i] = cat
	}

	return spyCatsList, nil
}

func (s *SpyCatService) Update(ctx context.Context, id uuid.UUID, args *UpdateSpyCatArgs) error {
	spyCat, err := s.spyCatRepository.GetById(ctx, id)
	if err != nil {
		return err
	}

	spyCat.Update(args.Salary)

	return s.spyCatRepository.Update(ctx, spyCat)
}
