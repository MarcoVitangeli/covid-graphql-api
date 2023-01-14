package cases

import (
	"context"
	"database/sql"
	"strings"
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

func (s *Service) Find(ctx context.Context, query *CaseSearch) ([]*Case, error) {
	var (
		cases []*Case
		qOps  []string
		qArgs []any
	)

	if query.Age != nil {
		qOps = append(qOps, "(age = ?)")
		qArgs = append(qArgs, *query.Age)
	}
	if query.Stage != nil {
		qOps = append(qOps, "(stage = ?)")
		qArgs = append(qArgs, *query.Stage)
	}
	if query.Dead != nil {
		qOps = append(qOps, "(dead = ?)")
		qArgs = append(qArgs, *query.Dead)
	}
	if query.Gender != nil {
		qOps = append(qOps, "(gender = ?)")
		qArgs = append(qArgs, *query.Gender)
	}
	if query.Province != nil {
		qOps = append(qOps, "(province = ?)")
		qArgs = append(qArgs, *query.Province)
	}
	if query.Neighborhood != nil {
		qOps = append(qOps, "(neighborhood = ?)")
		qArgs = append(qArgs, *query.Neighborhood)
	}
	q := `SELECT id, province, gender, neighborhood, age, stage, dead
FROM cases`

	if len(qOps) != 0 {
		q += "\nWHERE " + strings.Join(qOps, " AND ")
	}
	rows, err := s.DB.QueryContext(ctx, q, qArgs...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var c Case
		err := rows.Scan(&c.ID, &c.Province, &c.Gender, &c.Neighborhood, &c.Age, &c.Stage, &c.Dead)
		if err != nil {
			return nil, err
		}
		cases = append(cases, &c)
	}
	return cases, nil
}
