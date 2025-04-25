package repository

import "time"

type TestRepository struct{}

func NewTestRepository() *TestRepository {
	return &TestRepository{}
}

func (repo *TestRepository) Migrate() error {
	return nil
}

func (repo *TestRepository) InsertHolding(h Holdings) (*Holdings, error) {
	return &h, nil
}

func (repo *TestRepository) AllHoldings() ([]Holdings, error) {
	holdings := []Holdings{}
	return holdings, nil
}

func (repo *TestRepository) GetHoldingByID(id int) (*Holdings, error) {
	h := Holdings{
		Amount:        1,
		PurchaseDate:  time.Now(),
		PurchasePrice: 1000,
	}

	return &h, nil
}

func (repo *TestRepository) UpdateHolding(id int64, update Holdings) error {
	return nil
}

func (repo *TestRepository) DeleteHolding(id int64) error {
	return nil
}
