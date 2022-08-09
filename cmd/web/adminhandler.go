package main

import (
	"customerservice/cmd/web/models"
	"customerservice/internal/entity"
	"encoding/json"
	"errors"
	"github.com/kataras/jwt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func (app *App) adminLogin(w http.ResponseWriter, r *http.Request) {

	var request models.AdminLoginRequest
	var user entity.BackendUsers

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Query in db about the username

	tx := app.BackendUsers.Db.Where("username = ?", request.Username).First(&user)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		//
		http.Error(w, "User Not found", http.StatusNotFound)

	} else {
		// if the user is found check for the password as well
		if user.Password == request.Password {
			//Generate jwt and send

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
			http.Error(w, "Invalid Password", http.StatusUnauthorized)

		}

	}

}
