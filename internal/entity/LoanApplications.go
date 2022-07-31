package entity

import "time"

type LoanApplication struct {
	FirstName       string    `json:"first_name" gorm:"primaryKey" `
	SurName         string    `json:"sur_name"`
	MiddleName      string    `json:"middle_name"`
	Gender          string    `json:"gender"`
	PanNumber       string    `json:"pan_number"`
	Birthday        string    `json:"birthday"`
	CreatedAt       time.Time `json:"created_at"`
	UserId          int       `json:"user_id"` // this id will come from the users table
	LoanNumber      string    `json:"loan_number"`
	DisbusrseMentId int       `json:"disbusrsement_id"`
	PanCardImage    string    `json:"pancard_image"`
}
