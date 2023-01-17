package dataset

import "database/sql"

type Service struct {
	DB *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (d *Service) LoadDataset(ls []string) {
	//TODO: actually load the dataset
}
