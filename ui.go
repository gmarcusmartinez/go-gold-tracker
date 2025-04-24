package main

import "fyne.io/fyne/v2/container"

func (app *Config) makeUI() {
	/* PRICE TEXT */
	openPrice, currentPrice, priceChange := app.getPriceText()

	priceContent := container.NewGridWithColumns(3,
		openPrice,
		currentPrice,
		priceChange,
	)

	app.PriceContainer = priceContent

	toolbar := app.getToolbar()
	app.Toolbar = toolbar

	content := container.NewVBox(priceContent, toolbar)

	app.MainWindow.SetContent(content)

}
