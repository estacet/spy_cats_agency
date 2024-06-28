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

func NewTarget(missionId uuid.UUID, name string, country string) *Target {
	return &Target{
		Id:        uuid.New(),
		Name:      name,
		Country:   country,
		Status:    Started,
		MissionId: missionId,
	}
}

func (t *Target) UpdateNotes(notes string) error {
	if t.Mission.Status == Completed || t.Status == Completed {
		return errors.New("cannot update notes for Completed instance")
	}

	t.Notes = notes

	return nil
}

func (t *Target) Complete() {
	t.Status = Completed
}
