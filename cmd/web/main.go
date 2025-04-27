package main

import (
	"log"
	"net/http"
	"os"

	"github.com/iyilmaz24/Go-Id-Auth-Server/internal/config"
	"github.com/iyilmaz24/Go-Id-Auth-Server/internal/database"
	"github.com/iyilmaz24/Go-Id-Auth-Server/internal/database/models"
	_ "github.com/lib/pq"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	surveys  *models.SurveyModel
}

func main() {

	infoLog := log.New(os.Stdout, "***INFO LOG:\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "***ERROR LOG:\t", log.Ldate|log.Ltime|log.Lshortfile)

	appConfig := config.LoadConfig()

	db, err := database.OpenDB(appConfig.DbDsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		surveys:  &models.SurveyModel{DB: db},
	}

	srv := &http.Server{
		Addr:     appConfig.Port,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %v", srv.Addr)

	err = srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}
