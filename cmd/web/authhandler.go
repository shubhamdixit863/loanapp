package main

import (
	"customerservice/cmd/web/models"
	"customerservice/internal/entity"
	"customerservice/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/kataras/jwt"
	"net/http"
	"time"
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

	_, err = json.Marshal(&user)
	if err != nil {
		// handle error
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Checking for password of the user

	check := utils.CheckPasswordHash(request.Password, user.Password)

	if check {

		// generate jwt
		var sharedKey = []byte("sercrethatmaycontainch@r$32chars")

		token, err := jwt.Sign(jwt.HS256, sharedKey, user, jwt.MaxAge(15*time.Minute))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		js, err := json.Marshal(models.Response{
			Message: "Success",
			Status:  http.StatusOK,
			Data:    token,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

}

func (app *App) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var request models.SignupRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Hashing the password --before storing in the db

	hash, err := utils.HashPassword(request.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	user := entity.User{FirstName: request.FirstName, LastName: request.LastName, Phone: request.Phone, Email: request.Email, Password: hash}

	result := app.UserModel.Db.Create(&user) // pass pointer of data to Create
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	jData, err := json.Marshal(&user)
	if err != nil {
		// handle error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	w.Write(jData)

}
