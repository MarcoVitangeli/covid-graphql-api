package dataloader

import (
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

func (d *Service) LoadDataset(ls []string) error {
	//TODO: actually load the dataset
	for _, l := range ls {
		fields := strings.Split(l, ";")
		if len(fields) != 14 {
			return fmt.Errorf("error: expected 14 fields, %d found", len(fields))
		}
	}
	return nil
}
