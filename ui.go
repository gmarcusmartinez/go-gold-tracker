package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
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
			canvas.NewText("Holdings content goes here", nil)),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	content := container.NewVBox(priceContent, toolbar, tabs)

	app.MainWindow.SetContent(content)

}

func (app *Config) refreshPriceContent() {
	open, current, change := app.getPriceText()
	app.PriceContainer.Objects = []fyne.CanvasObject{open, current, change}
	app.PriceContainer.Refresh()

	chart := app.getChart()
	app.PriceChartContainer.Objects = []fyne.CanvasObject{chart}
	app.PriceChartContainer.Refresh()
}
