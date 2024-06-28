package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"spy-cats/internal/model"
	"spy-cats/pkg/apperror"
)

type MissionRepository struct {
	conn *pgx.Conn
}

func NewMissionRepository(conn *pgx.Conn) *MissionRepository {
	return &MissionRepository{conn: conn}
}

func (r *MissionRepository) GetById(ctx context.Context, id uuid.UUID) (*model.Mission, error) {
	query := `SELECT * FROM missions WHERE id = $1;`

	row := r.conn.QueryRow(ctx, query, id)

	mission := new(model.Mission)

	if err := row.Scan(
		&mission.ID,
		&mission.CatId,
		&mission.Status,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperror.NewEntityNotFoundError("Entity mission with id " + id.String() + " not found")
		}

		return nil, err
	}

	return mission, nil
}

func (r *MissionRepository) GetList(ctx context.Context) ([]*model.Mission, error) {
	query := `SELECT * FROM missions`

	rows, err := r.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var missions []*model.Mission

	for rows.Next() {
		mission := new(model.Mission)

		err := rows.Scan(
			&mission.ID,
			&mission.CatId,
			&mission.Status,
		)
		if err != nil {
			return nil, err
		}

		missions = append(missions, mission)
	}

	return missions, nil
}

func (r *MissionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM missions WHERE id = $1`

	_, err := r.conn.Exec(ctx, query, id)

	return err
}
