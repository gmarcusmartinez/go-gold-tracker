package main

import (
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
			canvas.NewText("Price content goes here", nil)),

		container.NewTabItemWithIcon(
			"Holdings", theme.InfoIcon(),
			canvas.NewText("Holdings content goes here", nil)),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	content := container.NewVBox(priceContent, toolbar, tabs)

	app.MainWindow.SetContent(content)

}
