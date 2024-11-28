package database

import (
	"database/sql"
	"time"
)

func OpenDB(dsn string) (*sql.DB, error) {

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	// db.SetMaxOpenConns(10) - don't use, creates artifical bottleneck of 10 connections
	db.SetMaxIdleConns(5) 		// keeps up to 5 idle connections at a time on standby
	db.SetConnMaxLifetime(10 * time.Minute) 		// recycles connections every 10 minutes
	
	return db, nil
}

