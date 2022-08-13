package main

import (
	"customerservice/cmd/web/models"
	"customerservice/internal/entity"
	"customerservice/internal/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func (app *App) UploadUserFormData(w http.ResponseWriter, r *http.Request) {
	var header = r.Header.Get("token")

	claims, err := utils.DecodeJwt(header)
	atoi, err := strconv.Atoi(claims["id"].(string))
	if err != nil {
		http.Error(w, "Jwt Parse Error", http.StatusBadRequest)
		return
	}

	if err != nil {
		log.Println(err)
	}
	fmt.Println(claims)
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	parent := filepath.Dir(filepath.Dir(path))
	fmt.Println(parent)

	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	uuidWithHyphen := uuid.New()

	uuidGenerated := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	fileExtension := strings.Split(handler.Filename, ".")[1]

	fileName := fmt.Sprintf("%s%s%s", parent, "/uploads/", uuidGenerated+"."+fileExtension)

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	//inserting the user data to the database
	loanNumber := utils.RandStringBytes(20)

	//Splitting the file name

	user := entity.LoanApplication{SurName: r.FormValue("sur_name"), FirstName: r.FormValue("first_name"),
		MiddleName: r.FormValue("middle_name"), Birthday: r.FormValue("birthday"), PanNumber: r.FormValue("pan_number"),
		Gender: r.FormValue("gender"), PancardImage: fmt.Sprintf("%s%s", "uploads/", strings.Split(fileName, "uploads/")[1]), LoanNumber: loanNumber,
		UserId: atoi,
		Status: "applied",
	}

	result := app.LoanApp.Db.Create(&user) // pass pointer of data to Create
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(models.Response{
		Message: "Success",
		Status:  http.StatusOK,
		Data:    user.Id,
	})

	if err != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	w.Write(j)

}

func (app *App) CheckLoanStatusHandler(w http.ResponseWriter, r *http.Request) {
	var header = r.Header.Get("token")
	var loanApplication entity.LoanApplication
	status := 2

	claims, err := utils.DecodeJwt(header)
	atoi, err := strconv.Atoi(claims["id"].(string))
	if err != nil {
		http.Error(w, "Jwt Parse Error", http.StatusBadRequest)
		return
	}
	tx := app.LoanApp.Db.First(&loanApplication, "id = ?", atoi)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		//
		status = 0

	} else if loanApplication.Status == "applied" {

		status = 1
	}

	js, _ := json.Marshal(models.Response{
		Message: "Success",
		Status:  200,
		Data:    status,
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func (app *App) ContactUploadHandler(w http.ResponseWriter, r *http.Request) {
	var request models.UploadContacts

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usercontact := entity.UserContact{
		Contacts: request.Contacts,
		LoanId:   request.LoanId,
	}

	_ = app.UserContactModel.Db.Create(&usercontact)

	// Updating the LoanApp table status to pending as well

	// Update with conditions
	app.LoanApp.Db.Model(&entity.LoanApplication{}).Where("id = ?", request.LoanId).Update("status", "pending")

	js, _ := json.Marshal(models.Response{
		Message: "Success",
		Status:  200,
		Data:    "Loan Application Success",
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
