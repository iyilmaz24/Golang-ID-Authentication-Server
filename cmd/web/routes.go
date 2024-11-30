package main

import "net/http"

func (app *application) routes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/getSurvey", app.getSurvey)
	mux.HandleFunc("/getSurveyDbHealth", app.getSurveyDbHealth)

	handler := app.enableCORS(mux)

	return handler
}