package main

import (
	"customerservice/cmd/web/models"
	"customerservice/internal/entity"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func (app *App) UploadUserFormData(w http.ResponseWriter, r *http.Request) {

	log.Println()

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

	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	fileExtension := strings.Split(handler.Filename, ".")[1]

	fileName := fmt.Sprintf("%s%s%s", parent, "/uploads/", uuid+"."+fileExtension)

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	//inserting the user data to the database

	user := entity.LoanApplication{SurName: r.FormValue("sur_name"), FirstName: r.FormValue("first_name"),
		MiddleName: r.FormValue("middle_name"), Birthday: r.FormValue("birthday"), PanNumber: r.FormValue("pan_number"),
		Gender: r.FormValue("gender"), PanCardImage: fileName,
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
		Data:    "SuccessFully Uploaded",
	})

	if err != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	w.Write(j)

}
