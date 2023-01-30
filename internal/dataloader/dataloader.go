package dataloader

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/MarcoVitangeli/covid-graphql-api/internal/logger"
	. "github.com/ahmetb/go-linq/v3"
	"go.uber.org/zap"
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

func (d *Service) truncateTable(ctx context.Context) error {
	_, err := d.DB.ExecContext(ctx, "TRUNCATE TABLE cases;")
	if err != nil {
		return err
	}
	return nil
}

func (d *Service) LoadDataset(ctx context.Context, ls []string) error {
	err := d.truncateTable(ctx)
	if err != nil {
		return err
	}

	cursor := 0
	stepSize := 100
	tx, err := d.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := From(ls)

	for {
		var lss []string
		query.Skip(cursor).Take(stepSize).ToSlice(&lss)
		if len(lss) == 0 {
			logger.Info("dataset processed succesfully")
			tx.Commit()
			return nil
		}
		err = loadDatasetLines(ctx, lss, tx)
		if err != nil {
			return err
		}
		cursor += stepSize
	}
}

func loadDatasetLines(ctx context.Context, ls []string, tx *sql.Tx) error {
	//TODO: actually load the dataset
	q := `INSERT INTO cases (id, province, gender, neighborhood, age, stage, dead) VALUES `
	vals := make([]string, 0, len(ls))
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
		// we do not process data with no age for convension
		if fs[indexAge] == "NA" {
			continue
		}

		vals = append(vals, "(?,?,?,?,?,?,?)")

		argsArr = append(argsArr,
			fs[indexId],
			fs[indexProvince],
			fs[indexGender],
			fs[indexNeighborhood],
			fs[indexAge],
			fs[indexStage],
			dead,
		)
	}

	logger.Info("data to process", zap.Int("amount", len(vals)))

	q += strings.Join(vals, ",")

	stmt, err := tx.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, argsArr...)
	return err
}
