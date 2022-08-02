package main

import (
	"customerservice/cmd/web/models"
	"customerservice/internal/entity"
	"customerservice/internal/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kataras/jwt"
	"gorm.io/gorm"
	"log"
	"net/http"
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
		v := "2014-05-03 20:57 UTC"
		then, err := time.Parse(timeFormat, v)
		if err != nil {
			fmt.Println(err)
			return
		}
		duration := time.Since(then)
		fmt.Println(duration.Minutes())

		//Check for otp as well
		if request.Otp == user.Otp && duration.Minutes() < 1 {
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
	_, err = utils.SendOtp(fmt.Sprintf("Your OTP Is %s", otp), request.Phone)

	if err != nil {

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

	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	js, _ := json.Marshal(models.Response{
		Message: "Success",
		Status:  200,
		Data:    "OTP Sent",
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
