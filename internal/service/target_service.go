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
	Notes  string       `json:"notes"`
	Status model.Status `json:"status"`
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

func (s *TargetService) Create(ctx context.Context, args *CreateTargetArgs) (*uuid.UUID, error) {
	mission, err := s.missionRepository.GetById(ctx, args.MissionId)
	if err != nil {
		return nil, err
	}

	if mission.Status == model.Completed {
		return nil, errors.New("cannot add target to completed mission")
	}

	target := model.NewTarget(
		args.MissionId,
		args.Name,
		args.Country,
		args.Notes,
	)

	if err := s.repository.Create(ctx, target); err != nil {
		return nil, err
	}

	return &target.Id, nil
}

func (s *TargetService) Update(ctx context.Context, id uuid.UUID, args *UpdateTargetArgs) error {
	target, err := s.repository.GetById(ctx, id)
	if err != nil {
		return err
	}

	mission, err := s.missionRepository.GetById(ctx, target.MissionId)
	if err != nil {
		return err
	}

	target.Mission = mission

	err = target.Update(args.Notes, args.Status)
	if err != nil {
		return err
	}

	return s.repository.Update(ctx, target)
}
