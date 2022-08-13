package main

import (
	"customerservice/cmd/web/models"
	"customerservice/internal/entity"
	"customerservice/internal/utils"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
)

//this check the otp
const (
	// See http://golang.org/pkg/time/#Parse
	timeFormat = "2006-01-02 15:04 MST"
)

func (app *App) LoginHAndler(w http.ResponseWriter, r *http.Request) {

	var user entity.MessageAuth
	var request models.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tx := app.MessageAuthModel.Db.Where("phone = ?", request.Phone).First(&user)

	if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		/*
			then, err := time.Parse(timeFormat, user.LastLogin.String())
			if err != nil {
				fmt.Println(err)
				return
			}
		*/
		duration := time.Since(user.LastLogin)
		fmt.Println(duration.Minutes())

		//Check for otp as well
		if request.Otp == user.Otp && duration.Minutes() < 10 {
			token, err := utils.GenerateToken(strconv.Itoa(user.Id))
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

		} else {
			http.Error(w, "OTP Didn't Match or Expired", http.StatusBadRequest)
			return
		}

	} else {
		http.Error(w, "User Not Found", http.StatusBadRequest)
		return
	}

}

// This send the otp to the user
func (app *App) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var request models.SignupRequest
	var user entity.MessageAuth
	//Decoding the json body in Login Request struct

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Send otp here --
	otp, _ := utils.GenerateOTP(5)

	_, err = utils.SendSMS(request.Phone, fmt.Sprintf("Your CashLo OTP is %s", otp))

	if err != nil {

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return

	} else {

		// works because destination struct is passed in
		tx := app.MessageAuthModel.Db.Where("phone = ?", request.Phone).First(&user).Error
		fmt.Println(tx)
		if errors.Is(tx, gorm.ErrRecordNotFound) {

			//Create a record in the db and send otp
			messageauth := entity.MessageAuth{CreatedAt: time.Now(), LastLogin: time.Now(), Phone: request.Phone, Ip: utils.GetIP(r), Otp: otp}

			result := app.MessageAuthModel.Db.Create(&messageauth) // pass pointer of data to Creat
			log.Println(result)

		} else {
			// just send  the otp and update the table
			// Update with conditions
			app.MessageAuthModel.Db.Model(&entity.MessageAuth{}).Where("phone = ?", request.Phone).Updates(&entity.MessageAuth{Otp: otp, LastLogin: time.Now()})
		}

	}
	js, _ := json.Marshal(models.Response{
		Message: "Success",
		Status:  200,
		Data:    "OTP Sent",
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
