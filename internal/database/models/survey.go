package models

import (
	"database/sql"
	"errors"

	"github.com/iyilmaz24/Go-Id-Auth-Server/internal/database"
)

type Survey struct {
	SurveyID   string
	Link       string
}

type SurveyModel struct {
	DB *sql.DB
}

type HealthCheck struct {
	Status          string `json:"status"`
	OpenConnections int    `json:"open_connections"`
	InUse           int    `json:"in_use"`
	Idle            int    `json:"idle"`
}

func (sm *SurveyModel) Get(id string) (*Survey, error) {
	sqlQuery := database.GetSurveyByIdQuery()
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

func (m *SurveyModel) CheckHealth() (*HealthCheck, error) {
	err := m.DB.Ping()
	status := "healthy"
	if err != nil {
		status = "unhealthy"
	}

	stats := m.DB.Stats()

	return &HealthCheck{
		Status:          status,
		OpenConnections: stats.OpenConnections,
		InUse:           stats.InUse,
		Idle:            stats.Idle,
	}, err
}

