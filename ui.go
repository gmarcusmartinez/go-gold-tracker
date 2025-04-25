package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func (app *Config) makeUI() {
	/* PRICE TEXT */
	openPrice, currentPrice, priceChange := app.getPriceText()

	priceContent := container.NewGridWithColumns(3,
		openPrice,
		currentPrice,
		priceChange,
	)

	holdingsTabContent := app.holdingsTab()
	app.PriceContainer = priceContent

	/* TOOLBAR */
	toolbar := app.getToolbar()
	app.Toolbar = toolbar

	/* APP TABS */
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon(
			"Prices",
			theme.HomeIcon(),
			app.pricesTab(),
		),
		container.NewTabItemWithIcon(
			"Holdings", theme.InfoIcon(),
			holdingsTabContent,
		),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	content := container.NewVBox(priceContent, toolbar, tabs)

	app.MainWindow.SetContent(content)

	/* REFRESH PRICE IN BACKGROUND EVERY 30 SECONDS */
	go func() {
		for range time.Tick(time.Second * 5) {
			app.refreshPriceContent()
		}
	}()

}

func (app *Config) refreshPriceContent() {
	open, current, change := app.getPriceText()

	fyne.Do(func() {
		app.PriceContainer.Objects = []fyne.CanvasObject{open, current, change}
		app.PriceContainer.Refresh()

		chart := app.getChart()
		app.PriceChartContainer.Objects = []fyne.CanvasObject{chart}
		app.PriceChartContainer.Refresh()
	})
}

func (app *Config) refreshHoldingsTable() {
	app.Holdings = app.getHoldingSlice()
	app.HoldingsTable.Refresh()
}
