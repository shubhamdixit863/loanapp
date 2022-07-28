package main

import (
	"customerservice/cmd/web/models"
	"customerservice/internal/entity"
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *App) LoginHAndler(w http.ResponseWriter, r *http.Request) {
	var request models.LoginRequest
	var user entity.User
	//Decoding the json body in Login Request struct

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// works because destination struct is passed in
	app.UserModel.Db.Where("email = ?", request.Email).Find(&user)
	fmt.Println(user)

	jData, err := json.Marshal(&user)
	if err != nil {
		// handle error
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
	//fmt.Fprintf(w, "Person: %+v", user)
}

func (app *App) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var request models.SignupRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := entity.User{FirstName: request.FirstName, LastName: request.LastName, Phone: request.Phone, Email: request.Email, Password: request.Password}

	result := app.UserModel.Db.Create(&user) // pass pointer of data to Create
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	jData, err := json.Marshal(&user)
	if err != nil {
		// handle error
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)

}
