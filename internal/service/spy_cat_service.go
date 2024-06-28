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
	Breed             string  `json:"breed"`
	Salary            float64 `json:"salary"`
}

type SpyCatDetails struct {
	Id                uuid.UUID `json:"id"`
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

//func (s *RaceService) GetList(ctx context.Context) ([]*RacesListItem, error) {
//	races, err := s.raceRepository.GetList(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	racesList := make([]*RacesListItem, len(races))
//
//	for i, race := range races {
//		raceItem := RacesListItem{Id: race.Id, Title: race.Title}
//		raceItemPointer := &raceItem
//		racesList[i] = raceItemPointer
//	}
//
//	return racesList, nil
//}

//func (s *RaceService) Update(ctx context.Context, id uuid.UUID, args *RaceArgs) error {
//	race, err := s.raceRepository.GetById(ctx, id)
//	if err != nil {
//		return err
//	}
//
//	race.Update(
//		args.Title,
//		args.Description,
//		args.MaxParticipantsCount,
//		args.RegistrationAt,
//		args.StartAt,
//		args.Category,
//	)
//
//	return s.raceRepository.Update(ctx, race)
//}
