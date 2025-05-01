package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/gmarcusmartinez/gold-tracker/repository"
)

func (app *Config) holdingsTab() *fyne.Container {
	app.Holdings = app.getHoldingSlice()
	app.HoldingsTable = app.buildHoldingsTable()

	holdingsContainer := container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		container.NewAdaptiveGrid(1, app.HoldingsTable),
	)

	return holdingsContainer

}

func (app *Config) buildHoldingsTable() *widget.Table {
	t := widget.NewTable(
		func() (int, int) {
			return len(app.Holdings), len(app.Holdings[0])
		},
		func() fyne.CanvasObject {
			ctr := container.NewVBox(widget.NewLabel(""))
			return ctr
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Col == (len(app.Holdings[0])-1) && i.Row != 0 {
				w := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(),
					func() {
						app.showConfirmDeleteDialog(app.Holdings, i)
					})
				w.Importance = widget.HighImportance
				o.(*fyne.Container).Objects = []fyne.CanvasObject{w}
			} else {
				o.(*fyne.Container).Objects = []fyne.CanvasObject{
					widget.NewLabel(app.Holdings[i.Row][i.Col].(string)),
				}
			}
		},
	)

	colWidths := []float32{50, 200, 200, 200, 110}
	for i := range colWidths {
		t.SetColumnWidth(i, colWidths[i])
	}

	return t
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
		currentAmount := float32(entry.PurchasePrice) / 100
		currentRow = append(currentRow, fmt.Sprintf("$%2f ", currentAmount))
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

func (app *Config) showConfirmDeleteDialog(data [][]interface{}, i widget.TableCellID) {
	dialog.ShowConfirm("Delete", "", func(deleted bool) {
		if deleted {
			id, _ := strconv.Atoi(data[i.Row][0].(string))

			err := app.DB.DeleteHolding(int64(id))
			if err != nil {
				app.ErrorLog.Println(err)
			}
		}

		app.refreshHoldingsTable()
	}, app.MainWindow)
}
