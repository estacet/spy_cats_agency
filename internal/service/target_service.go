package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"spy-cats/internal/model"
	"spy-cats/internal/repository"
)

type CreateTargetArgs struct {
	MissionId uuid.UUID    `json:"mission_id"`
	Name      string       `json:"name"`
	Country   string       `json:"country"`
	Notes     string       `json:"notes"`
	Status    model.Status `json:"status"`
}

type UpdateTargetArgs struct {
	Notes string `json:"notes"`
}

type TargetService struct {
	repository        *repository.TargetRepository
	missionRepository *repository.MissionRepository
}

func NewTargetService(
	repository *repository.TargetRepository,
	missionRepository *repository.MissionRepository,
) *TargetService {
	return &TargetService{
		repository:        repository,
		missionRepository: missionRepository,
	}
}

func (s *TargetService) Create(ctx context.Context, args *CreateTargetArgs) error {
	mission, err := s.missionRepository.GetById(ctx, args.MissionId)
	if err != nil {
		return err
	}

	if mission.Status == model.Completed {
		return errors.New("cannot add target to completed mission")
	}

	target := model.NewTarget(
		args.MissionId,
		args.Name,
		args.Country,
		args.Notes,
	)

	if err := s.repository.Create(ctx, target); err != nil {
		return err
	}

	return nil
}
