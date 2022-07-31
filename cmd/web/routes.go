package main

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func (app *App) registerroutes() *mux.Router {
	r := mux.NewRouter()

	r.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	// documentation for developers
	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)
	r.Handle("/docs", sh)

	// documentation for share
	opts1 := middleware.RedocOpts{SpecURL: "/swagger.yaml", Path: "docs1"}
	sh1 := middleware.Redoc(opts1, nil)
	r.Handle("/docs1", sh1)

	r.HandleFunc("/v1/health", app.HealthcheckHandler).Methods("GET")
	r.HandleFunc("/v1/login", app.LoginHAndler).Methods("POST")
	r.HandleFunc("/v1/signup", app.SignupHandler).Methods("POST")
	r.HandleFunc("/v1/formdata", app.UploadUserFormData).Methods("POST")

	return r

}
