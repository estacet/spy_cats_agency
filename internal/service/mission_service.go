package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"spy-cats/internal/model"
	"spy-cats/internal/repository"
)

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
}

func NewMissionService(missionRepository *repository.MissionRepository) *MissionService {
	return &MissionService{missionRepository: missionRepository}
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
