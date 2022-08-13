package entity

import (
	"gorm.io/gorm"
	"time"
)

type UserContact struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	LoanId    int       `json:"loan_id"`
	Contacts  string    `json:"contacts"`
	CreatedAt time.Time `json:"created_at"`
}

func (UserContact) TableName() string {
	return "user_contacts"
}

type UserContactModel struct {
	Db *gorm.DB
}
