package model

import "github.com/google/uuid"

type Status string

const (
	Initiated Status = "initiated"
	Ongoing   Status = "ongoing"
	Completed Status = "completed"
)

type Mission struct {
	ID     uuid.UUID
	CatId  uuid.UUID
	Status Status
}

func NewMission(catId uuid.UUID) *Mission {
	return &Mission{
		ID:     uuid.New(),
		CatId:  catId,
		Status: Initiated,
	}
}

func (m *Mission) UpdateMissionStatus(newStatus Status) {
	m.Status = newStatus
}
