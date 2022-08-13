package main

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func init() {

	godotenv.Load()

}

func (app *App) registerroutes() *mux.Router {
	r := mux.NewRouter()
	var specurl string = "/swagger.yaml"

	if os.Getenv("ENV") == "local" {
		specurl = "/swagger.local.yaml"
	}

	r.Handle(specurl, http.FileServer(http.Dir("./")))
	// documentation for developers
	opts := middleware.SwaggerUIOpts{SpecURL: specurl}
	sh := middleware.SwaggerUI(opts, nil)
	r.Handle("/docs", sh)

	// documentation for share
	opts1 := middleware.RedocOpts{SpecURL: specurl, Path: "docs1"}
	sh1 := middleware.Redoc(opts1, nil)
	r.Handle("/docs1", sh1)

	//Serving static content
	fs := http.FileServer(http.Dir("../../uploads/"))
	r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", fs))

	r.HandleFunc("/v1/health", app.HealthcheckHandler).Methods("GET")
	r.HandleFunc("/v1/login", app.LoginHAndler).Methods("POST")
	r.HandleFunc("/v1/signup", app.SignupHandler).Methods("POST")
	formdata := r.PathPrefix("/v1/verified").Subrouter()
	formdata.Use(JwtVerify)
	formdata.HandleFunc("/formdata", app.UploadUserFormData).Methods("POST")
	formdata.HandleFunc("/contactupload", app.ContactUploadHandler).Methods("POST")

	adminRoutes := r.PathPrefix("/v1/admin").Subrouter()

	adminRoutes.HandleFunc("/login", app.adminLogin).Methods("POST")
	adminRoutes.HandleFunc("/loandata", app.GetLoanApplicationsAdmin).Methods("GET")

	return r

}
