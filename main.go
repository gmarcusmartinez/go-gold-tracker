package main

import (
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

type Config struct {
	App            fyne.App
	InfoLog        *log.Logger
	ErrorLog       *log.Logger
	MainWindow     fyne.Window
	PriceContainer *fyne.Container
	Toolbar        *widget.Toolbar
	HTTPClient     *http.Client
}

var cfg Config

func main() {
	a := app.NewWithID("bd.gocode.goldtracker.preferences")
	cfg.App = a

	cfg.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	cfg.ErrorLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
	cfg.HTTPClient = &http.Client{}

	// TODO: open a connection to the database

	// TODO: create a database repository

	cfg.MainWindow = a.NewWindow("GoldTracker")
	cfg.MainWindow.Resize(fyne.NewSize(770, 410))
	cfg.MainWindow.SetFixedSize(true)
	cfg.MainWindow.SetMaster()

	cfg.makeUI()

	cfg.MainWindow.ShowAndRun()

}
