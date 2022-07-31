package models

type LoginRequest struct {
	Email string `json:"email"`

	Password string `json:"password"`
}

type SignupRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
}

type FormDateRequest struct {
	FirstName  string `json:"first_name"`
	SurName    string `json:"sur_name"`
	MiddleName string `json:"middle_name"`
	Gender     string `json:"gender"`
	PanNumber  string `json:"pan_number"`
	Birthday   string `json:"birthday"`
}
