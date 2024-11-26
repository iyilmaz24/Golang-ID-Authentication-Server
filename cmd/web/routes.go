package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/getSurvey", app.getSurvey)
	mux.HandleFunc("/getSurveyDbHealth", app.getSurveyDbHealth)

	return mux
}