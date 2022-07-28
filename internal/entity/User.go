package entity

import "gorm.io/gorm"

type User struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
}

type UserModel struct {
	Db *gorm.DB
}
