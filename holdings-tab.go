package main

import (
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
