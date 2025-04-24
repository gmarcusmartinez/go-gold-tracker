package repository

import (
	"errors"
	"time"
)

var (
	errorUpdateFailed = errors.New("update failed")
	errorDeleteFailed = errors.New("delete failed")
)

type Repository interface {
	Migrate() error
	InsertHolding(h Holdings) (*Holdings, error)
	AllHoldings() ([]Holdings, error)
	GetHoldingByID(id int) (*Holdings, error)
	UpdateHolding(id int64, update Holdings) error
	DeleteHolding(id int64) error
}

type Holdings struct {
	ID            int64     `json:"id"`
	Amount        int       `json:"amount"`
	PurchaseDate  time.Time `json:"purchase_date"`
	PurchasePrice int       `json:"purchase_price"`
}
