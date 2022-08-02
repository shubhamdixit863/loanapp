package models

type SignupRequest struct {
	Phone string `json:"phone"`
}

type LoginRequest struct {
	Phone string `json:"phone"`
	Otp   string `json:"otp"`
}

type FormDateRequest struct {
	FirstName  string `json:"first_name"`
	SurName    string `json:"sur_name"`
	MiddleName string `json:"middle_name"`
	Gender     string `json:"gender"`
	PanNumber  string `json:"pan_number"`
	Birthday   string `json:"birthday"`
}
