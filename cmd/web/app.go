package main

import (
	"customerservice/internal/entity"
	"log"
)

type App struct {
	ErrorLog  *log.Logger
	InfoLog   *log.Logger
	UserModel *entity.UserModel
}
