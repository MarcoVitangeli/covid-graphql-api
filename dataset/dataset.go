package dataset

import "database/sql"

type DatasetService struct {
	DB *sql.DB
}

func NewService(db *sql.DB) *DatasetService {
	return &DatasetService{
		DB: db,
	}
}

func (d *DatasetService) LoadDataset(ls []string) {
	//TODO: actually load the dataset
}
