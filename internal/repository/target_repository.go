package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"spy-cats/internal/model"
	"spy-cats/pkg/apperror"
)

type TargetRepository struct {
	conn *pgx.Conn
}

func NewTargetRepository(conn *pgx.Conn) *TargetRepository {
	return &TargetRepository{conn: conn}
}

func (r *TargetRepository) Create(ctx context.Context, target *model.Target) error {
	query := `
INSERT INTO targets (id, name, country, notes, status, mission_id)
VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.conn.Exec(ctx, query,
		target.Id,
		target.Name,
		target.Country,
		target.Notes,
		target.Status,
		target.MissionId,
	)

	return err
}

func (r *TargetRepository) GetById(ctx context.Context, id uuid.UUID) (*model.Target, error) {
	query := `SELECT * FROM targets WHERE id = $1;`

	row := r.conn.QueryRow(ctx, query, id)

	target := new(model.Target)

	if err := row.Scan(
		&target.Id,
		&target.MissionId,
		&target.Name,
		&target.Country,
		&target.Notes,
		&target.Status,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperror.NewEntityNotFoundError("Entity target with id " + id.String() + " not found")
		}

		return nil, err
	}

	return target, nil
}

func (r *TargetRepository) Update(ctx context.Context, target *model.Target) error {
	query := `UPDATE targets
		SET (notes, status) = ($2, $3)
		WHERE id = $1;`

	_, err := r.conn.Exec(ctx, query,
		target.Id,
		target.Notes,
		target.Status,
	)

	return err
}
