package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"spy-cats/internal/model"
	"spy-cats/internal/repository"
)

type CreateMissionArgs struct {
	CatId   uuid.NullUUID   `json:"cat_id"`
	Targets []MissionTarget `json:"targets"`
}

type MissionTarget struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

type UpdateMissingArgs struct {
	CatId  uuid.NullUUID `json:"cat_id"`
	Status model.Status  `json:"status"`
}

type MissionDetails struct {
	ID     uuid.UUID     `json:"id"`
	CatId  uuid.NullUUID `json:"cat_id"`
	Status model.Status  `json:"status"`
}

type MissionListItem struct {
	ID     uuid.UUID    `json:"id"`
	Status model.Status `json:"status"`
}

type MissionService struct {
	missionRepository *repository.MissionRepository
	targetRepository  *repository.TargetRepository
}

func NewMissionService(
	missionRepository *repository.MissionRepository,
	targetRepository *repository.TargetRepository,
) *MissionService {
	return &MissionService{
		missionRepository: missionRepository,
		targetRepository:  targetRepository,
	}
}

func (s *MissionService) GetById(ctx context.Context, id uuid.UUID) (*MissionDetails, error) {
	mission, err := s.missionRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	missionDetails := &MissionDetails{
		ID:     mission.ID,
		CatId:  mission.CatId,
		Status: mission.Status,
	}

	return missionDetails, nil
}

func (s *MissionService) GetList(ctx context.Context) ([]*MissionListItem, error) {
	missions, err := s.missionRepository.GetList(ctx)
	if err != nil {
		return nil, err
	}

	missionList := make([]*MissionListItem, len(missions))

	for i, mission := range missions {
		missionItem := &MissionListItem{
			ID:     mission.ID,
			Status: mission.Status,
		}

		missionList[i] = missionItem
	}

	return missionList, nil
}

func (s *MissionService) Delete(ctx context.Context, id uuid.UUID) error {
	mission, err := s.missionRepository.GetById(ctx, id)
	if err != nil {
		return err
	}

	if mission.IsCatAssigned() {
		return errors.New("cat already assigned to this mission, so it cannot be deleted")
	}

	return s.missionRepository.Delete(ctx, id)
}

// TODO: wrap into database transaction
func (s *MissionService) Create(ctx context.Context, args *CreateMissionArgs) (*uuid.UUID, error) {
	mission := model.NewMission(args.CatId)

	if err := s.missionRepository.Create(ctx, mission); err != nil {
		return nil, err
	}

	for _, missionTarget := range args.Targets {
		target := model.NewTarget(mission.ID, missionTarget.Name, missionTarget.Country)
		if err := s.targetRepository.Create(ctx, target); err != nil {
			return nil, err
		}
	}

	return &mission.ID, nil
}

func (s *MissionService) Update(ctx context.Context, id uuid.UUID, args *UpdateMissingArgs) error {
	mission, err := s.missionRepository.GetById(ctx, id)
	if err != nil {
		return err
	}

	if args.CatId.Valid {
		mission.AssignCat(args.CatId.UUID)
	}

	if args.Status == model.Completed {
		mission.Complete()
	}

	return s.missionRepository.Update(ctx, mission)
}
