package entity

import (
	"gorm.io/gorm"
	"time"
)

type LoanApplication struct {
	Id             int       `json:"id" gorm:"primaryKey"`
	FirstName      string    `json:"first_name" gorm:"primaryKey" `
	SurName        string    `json:"sur_name"`
	MiddleName     string    `json:"middle_name"`
	Gender         string    `json:"gender"`
	PanNumber      string    `json:"pan_number"`
	Birthday       string    `json:"birthday"`
	CreatedAt      time.Time `json:"created_at"`
	UserId         int       `json:"user_id"` // this id will come from the users table
	LoanNumber     string    `json:"loan_number"`
	DisbursementId int       `json:"disbursement_id"`
	PancardImage   string    `json:"pancard_image"`
}

// TableName overrides the table name used by User to `profiles`
func (LoanApplication) TableName() string {
	return "loan_applications"
}

type LoanApplicationModel struct {
	Db *gorm.DB
}
