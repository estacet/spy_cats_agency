package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"spy-cats/internal/model"
	"spy-cats/pkg/apperror"
)

type SpyCatRepository struct {
	conn *pgx.Conn
}

func NewSpyCatRepository(conn *pgx.Conn) *SpyCatRepository {
	return &SpyCatRepository{conn: conn}
}

func (r *SpyCatRepository) Create(ctx context.Context, spyCat *model.SpyCat) error {
	query := `
INSERT INTO spy_cats (id, name, experience_years, breed, salary)
VALUES ($1, $2, $3, $4, $5)`

	_, err := r.conn.Exec(ctx, query,
		spyCat.ID,
		spyCat.Name,
		spyCat.YearsOfExperience,
		spyCat.Breed,
		spyCat.Salary,
	)

	return err
}

func (r *SpyCatRepository) GetById(ctx context.Context, id uuid.UUID) (*model.SpyCat, error) {
	query := `SELECT * FROM spy_cats WHERE id = $1;`

	row := r.conn.QueryRow(ctx, query, id)

	spyCat := new(model.SpyCat)

	if err := row.Scan(
		&spyCat.ID,
		&spyCat.Name,
		&spyCat.YearsOfExperience,
		&spyCat.Breed,
		&spyCat.Salary,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperror.NewEntityNotFoundError("Entity spyCat with id " + id.String() + " not found")
		}

		return nil, err
	}

	return spyCat, nil
}

func (r *SpyCatRepository) GetList(ctx context.Context) ([]*model.SpyCat, error) {
	query := `SELECT * FROM spy_cats`

	rows, err := r.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var spyCatsList []*model.SpyCat

	for rows.Next() {
		spyCat := new(model.SpyCat)

		err := rows.Scan(
			&spyCat.ID,
			&spyCat.Name,
			&spyCat.YearsOfExperience,
			&spyCat.Breed,
			&spyCat.Salary,
		)
		if err != nil {
			return nil, err
		}

		spyCatsList = append(spyCatsList, spyCat)
	}

	return spyCatsList, nil
}

func (r *SpyCatRepository) Update(ctx context.Context, spyCat *model.SpyCat) error {
	query := `UPDATE spy_cats
		SET salary = $2
		WHERE id = $1;`

	_, err := r.conn.Exec(ctx, query,
		spyCat.ID,
		spyCat.Salary,
	)

	return err
}
