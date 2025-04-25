package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/gmarcusmartinez/gold-tracker/repository"
)

func (app *Config) holdingsTab() *fyne.Container {
	return nil
}

func (app *Config) getHoldingsTable() *widget.Table {
	return nil
}

func (app *Config) getHoldingSlice() [][]interface{} {
	var slice [][]interface{}

	holdings, err := app.currentHoldings()
	if err != nil {
		app.ErrorLog.Println(err)
	}

	slice = append(slice, []interface{}{
		"ID", "Amount", "Price", "Date", "Delete",
	})

	for _, entry := range holdings {
		var currentRow []interface{}
		currentRow = append(currentRow, strconv.FormatInt(entry.ID, 10))
		currentRow = append(currentRow, fmt.Sprintf("%d toz", entry.Amount))
		currentRow = append(currentRow, fmt.Sprintf("$%2f ", float32(entry.PurchasePrice/100)))
		currentRow = append(currentRow, entry.PurchaseDate.Format("2006-01-02"))
		currentRow = append(currentRow, widget.NewButton("Delete", func() {}))

		slice = append(slice, currentRow)
	}

	return slice
}

func (app *Config) currentHoldings() ([]repository.Holdings, error) {
	holdings, err := app.DB.AllHoldings()
	if err != nil {
		app.ErrorLog.Println(err)
		return nil, err
	}

	return holdings, nil
}
