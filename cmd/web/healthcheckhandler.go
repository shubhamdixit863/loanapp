package main

import "net/http"

func (app *App) HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Apis Working  Fine My friend  All good Friend!!"))

}
