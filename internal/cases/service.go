package cases

import (
	"context"
	"database/sql"
)

type Service struct {
	DB *sql.DB //TODO: use a storage interface instead of raw sql pointer
}

func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) Get(ctx context.Context, id int) (Case, error) {
	var c Case

	query := `SELECT id, province, gender, neighborhood, age, stage, dead
FROM cases
WHERE id = ?`

	err := s.DB.QueryRowContext(ctx, query, id).
		Scan(&c.ID, &c.Province, &c.Gender, &c.Neighborhood, &c.Age, &c.Stage, &c.Dead)
	if err != nil {
		return Case{}, err
	}
	return c, nil
}
