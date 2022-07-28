package main

import (
	"github.com/gorilla/mux"
)

func (app *App) registerroutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/v1/health", app.HealthcheckHandler).Methods("GET")
	r.HandleFunc("/v1/login", app.LoginHAndler).Methods("POST")
	r.HandleFunc("/v1/signup", app.SignupHandler).Methods("POST")

	return r

}
