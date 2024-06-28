package model

import "github.com/google/uuid"

type Status string

const (
	Started   Status = "started"
	Completed Status = "completed"
)

type Mission struct {
	ID     uuid.UUID
	CatId  uuid.NullUUID
	Status Status
}

func NewMission(catId uuid.NullUUID) *Mission {
	return &Mission{
		ID:     uuid.New(),
		CatId:  catId,
		Status: Started,
	}
}

func (m *Mission) IsCatAssigned() bool {
	return m.CatId.Valid
}

func (m *Mission) AssignCat(catId uuid.UUID) {
	m.CatId.UUID = catId
}

func (m *Mission) Complete() {
	m.Status = Completed
}
