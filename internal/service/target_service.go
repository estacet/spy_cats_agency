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
	Status    model.Status `json:"status"`
}

type UpdateTargetArgs struct {
	Notes  string       `json:"notes"`
	Status model.Status `json:"status"`
}

type TargetService struct {
	targetRepository  *repository.TargetRepository
	missionRepository *repository.MissionRepository
}

func NewTargetService(
	targetRepository *repository.TargetRepository,
	missionRepository *repository.MissionRepository,
) *TargetService {
	return &TargetService{
		targetRepository:  targetRepository,
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
	)

	if err := s.targetRepository.Create(ctx, target); err != nil {
		return nil, err
	}

	return &target.Id, nil
}

func (s *TargetService) Update(ctx context.Context, id uuid.UUID, args *UpdateTargetArgs) error {
	target, err := s.targetRepository.GetById(ctx, id)
	if err != nil {
		return err
	}

	mission, err := s.missionRepository.GetById(ctx, target.MissionId)
	if err != nil {
		return err
	}

	target.Mission = mission

	err = target.UpdateNotes(args.Notes)
	if err != nil {
		return err
	}

	if args.Status == model.Completed {
		target.Complete()
	}

	return s.targetRepository.Update(ctx, target)
}

func (s *TargetService) Delete(ctx context.Context, id uuid.UUID) error {
	target, err := s.targetRepository.GetById(ctx, id)
	if err != nil {
		return err
	}

	if target.Status == model.Completed {
		return errors.New("target already completed, so cannot be deleted")
	}

	return s.targetRepository.Delete(ctx, id)
}
