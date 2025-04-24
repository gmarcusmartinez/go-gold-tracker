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
	content := container.NewVBox(priceContent)

	app.MainWindow.SetContent(content)

}
