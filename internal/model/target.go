package model

import (
	"errors"
	"github.com/google/uuid"
)

type Target struct {
	Id        uuid.UUID
	MissionId uuid.UUID
	Name      string
	Country   string
	Notes     string
	Status    Status

	Mission *Mission
}

func NewTarget(missionId uuid.UUID, name string, country string, notes string) *Target {
	return &Target{
		Id:        uuid.New(),
		Name:      name,
		Country:   country,
		Notes:     notes,
		Status:    Initiated,
		MissionId: missionId,
	}
}

func (t *Target) Update(notes string, status Status) error {
	if t.Mission.Status == Completed || t.Status == Completed {
		return errors.New("cannot update notes for Completed instance")
	}

	t.Notes = notes
	t.Status = status
	return nil
}
