package repository

import (
	"database/sql"
	"time"
)

type SQLiteRepository struct {
	Conn *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		Conn: db,
	}
}

func (repo *SQLiteRepository) Migrate() error {
	query := `
	CREATE TABLE IF NOT EXISTS holdings (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		amount REAL NOT NULL,
		purchase_date INTEGER NOT NULL,
		purchase_price INTEGER NOT NULL
	)
	`
	_, err := repo.Conn.Exec(query)
	return err
}

func (repo *SQLiteRepository) InsertHolding(h Holdings) (*Holdings, error) {
	query := `
	INSERT INTO holdings (amount, purchase_date, purchase_price)
	VALUES (?, ?, ?)
	`
	res, err := repo.Conn.Exec(query, h.Amount, h.PurchaseDate.Unix(), h.PurchasePrice)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	h.ID = id

	return &h, nil
}

func (repo *SQLiteRepository) AllHoldings() ([]Holdings, error) {
	query := `
	SELECT id, amount, purchase_date, purchase_price FROM holdings
	ORDER BY purchase_date DESC
	`
	rows, err := repo.Conn.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	holdings := []Holdings{}
	for rows.Next() {
		var h Holdings
		var unixTime int64

		err = rows.Scan(
			&h.ID,
			&h.Amount,
			&unixTime,
			&h.PurchasePrice,
		)
		if err != nil {
			return nil, err
		}

		h.PurchaseDate = time.Unix(unixTime, 0)
		holdings = append(holdings, h)
	}

	return holdings, nil
}

func (repo *SQLiteRepository) GetHoldingByID(id int) (*Holdings, error) {
	query := `
	SELECT id, amount, purchase_date, purchase_price FROM holdings
	WHERE id = ?
	`
	row := repo.Conn.QueryRow(query, id)

	var h Holdings
	var unixTime int64

	err := row.Scan(&h.ID, &h.Amount, &unixTime, &h.PurchasePrice)
	if err != nil {
		return nil, err
	}

	h.PurchaseDate = time.Unix(unixTime, 0)
	return &h, nil
}

func (repo *SQLiteRepository) UpdateHolding(id int64, update Holdings) error {
	query := `
	UPDATE holdings SET amount = ?, purchase_date = ?, purchase_price = ?
	WHERE id = ?
	`
	res, err := repo.Conn.Exec(query,
		update.Amount,
		update.PurchaseDate.Unix(),
		update.PurchasePrice,
		id,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errorUpdateFailed
	}

	return nil
}

func (repo *SQLiteRepository) DeleteHolding(id int64) error {
	query := `
	DELETE FROM holdings WHERE id = ?
	`
	res, err := repo.Conn.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errorDeleteFailed
	}

	return nil
}
