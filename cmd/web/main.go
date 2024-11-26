package main

import (
	"log"

	"github.com/iyilmaz24/Go-Id-Auth-Server/internal/database/models"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	surveys  *models.SurveyModel
}


