package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/gmarcusmartinez/gold-tracker/repository"

	_ "github.com/glebarez/go-sqlite"
)

type Config struct {
	App                 fyne.App
	InfoLog             *log.Logger
	ErrorLog            *log.Logger
	DB                  repository.Repository
	MainWindow          fyne.Window
	PriceContainer      *fyne.Container
	Toolbar             *widget.Toolbar
	PriceChartContainer *fyne.Container
	Holdings            [][]interface{}
	HoldingsTable       *widget.Table
	HTTPClient          *http.Client
}

func main() {
	var cfg Config

	a := app.NewWithID("bd.gocode.goldtracker.preferences")
	cfg.App = a

	cfg.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	cfg.ErrorLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
	cfg.HTTPClient = &http.Client{}

	/* CONNECT AND SETUP DB */
	db, err := cfg.connectDB()
	if err != nil {
		log.Panic(err)
	}
	cfg.setupDB(db)

	cfg.MainWindow = a.NewWindow("GoldTracker")
	cfg.MainWindow.Resize(fyne.NewSize(770, 410))
	cfg.MainWindow.SetFixedSize(true)
	cfg.MainWindow.SetMaster()

	cfg.makeUI()

	cfg.MainWindow.ShowAndRun()

}

func (app *Config) connectDB() (*sql.DB, error) {
	path := ""

	if os.Getenv("DB_PATH") != "" {
		path = os.Getenv("DB_PATH")
	} else {
		path = app.App.Storage().RootURI().Path() + "/sql.db"
		app.InfoLog.Println("No DB_PATH environment variable, using default path:", path)
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (app *Config) setupDB(sqlDB *sql.DB) {
	app.DB = repository.NewSQLiteRepository(sqlDB)
	err := app.DB.Migrate()
	if err != nil {
		app.ErrorLog.Println(err)
		log.Panic(err)
	}
}
