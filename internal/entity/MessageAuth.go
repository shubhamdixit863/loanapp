package entity

import (
	"gorm.io/gorm"
	"time"
)

type MessageAuth struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Otp       string    `json:"otp"`
	Ip        string    `json:"ip"`
	LastLogin time.Time `json:"last_login"`
	CreatedAt time.Time `json:"created_at"`
	Phone     string    `json:"phone"`
}

func (MessageAuth) TableName() string {
	return "message_auth"
}

type MessageAuthModel struct {
	Db *gorm.DB
}
