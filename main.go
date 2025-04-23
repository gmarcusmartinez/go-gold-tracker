package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Config struct {
	App      fyne.App
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

var cfg Config

func main() {
	a := app.NewWithID("bd.gocode.goldtracker.preferences")
	cfg.App = a

	cfg.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	cfg.ErrorLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)

	// TODO: open a connection to the database

	// TODO: create a database repository

	w := a.NewWindow("GoldTracker")
	w.ShowAndRun()

}
