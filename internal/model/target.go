package model

import "github.com/google/uuid"

type Target struct {
	Id        uuid.UUID
	MissionId uuid.UUID
	Name      string
	Country   string
	Notes     string
	Status    Status
}

func NewTarget(missionId uuid.UUID, name string, country string, notes string) *Target {
	return &Target{
		Id:        uuid.New(),
		MissionId: missionId,
		Name:      name,
		Country:   country,
		Notes:     notes,
		Status:    Initiated,
	}
}
