package entity

import (
	"gorm.io/gorm"
	"time"
)

type BackendUsers struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`

	LastLogin time.Time `json:"last_login"`
	CreatedAt time.Time `json:"created_at"`
}

// TableName overrides the table name used by User to `profiles`
func (BackendUsers) TableName() string {
	return "backend_users"
}

type BackendUsersModel struct {
	Db *gorm.DB
}
