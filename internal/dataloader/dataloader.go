package dataloader

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type Service struct {
	DB *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}

/**
TODO: este programa no compilar porque MySQL no soporta tantos placeholders (?)
en una query, por lo cual, debemos procesar estas queries por batches (sobre una misma
transaccion, ya que queremos que sea atomico)
*/

const (
	indexId           = 0
	indexProvince     = 4
	indexGender       = 7
	indexNeighborhood = 5
	indexAge          = 8
	indexStage        = 9
	indexDead         = 11
	fieldsLen         = 14
)

func (d *Service) LoadDataset(ctx context.Context, ls []string) error {
	//TODO: actually load the dataset
	q := `INSERT INTO cases (id, province, gender, neighborhood, age, stage, dead) VALUES `
	vals := make([]string, 0, len(ls))
	for i := 0; i < len(ls); i++ {
		vals = append(vals, "(?,?,?,?,?,?,?)")
	}
	q += strings.Join(vals, ",")
	argsArr := make([]any, 0, len(ls)*7)
	for _, l := range ls {
		fs := strings.Split(l, ";")
		if len(fs) != fieldsLen {
			return fmt.Errorf("length of fields is shorter than expected: expected 14, found %d", len(fs))
		}
		dead := "YES"
		if fs[indexDead] == "NA" {
			dead = "NO"
		}
		argsArr = append(argsArr,
			fs[indexId],
			fs[indexProvince],
			fs[indexGender],
			fs[indexNeighborhood],
			fs[indexAge],
			fs[indexStage],
			dead)
	}
	stmt, err := d.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, argsArr...)
	return err
}
