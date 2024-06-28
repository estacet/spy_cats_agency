package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"spy-cats/internal/model"
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
