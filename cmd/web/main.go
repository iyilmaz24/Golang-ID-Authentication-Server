package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/iyilmaz24/Go-Id-Auth-Server/internal/database/models"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	surveys  *models.SurveyModel
}

func main() {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	dsn, ok := os.LookupEnv("DB_DSN")
	if !ok {
		errorLog.Fatal("DSN not set in environment variables")
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = ":3000"
	}

	db, err := openDB(dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
		surveys: &models.SurveyModel{DB: db},
	}

	srv := &http.Server{
		Addr:     port,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %v", port)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
