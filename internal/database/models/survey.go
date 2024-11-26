package models

import (
	"database/sql"
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"
)

type Survey struct {
	SurveyID   string
	Link       string
}

type SurveyModel struct {
	DB *sql.DB
}

func LoadQuery (fileName string) string {
	path := filepath.Join("internal", "database", "sql", fileName)

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to load query %s: %v", fileName, err)
	}

	return string(content)
}

func (sm *SurveyModel) Get(id string) (*Survey, error) {
	sqlQuery := LoadQuery("get_survey_by_id.sql")
	row := sm.DB.QueryRow(sqlQuery, id)

	s := &Survey{}
	err := row.Scan(&s.Link, &s.SurveyID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}
	
	return s, nil
}