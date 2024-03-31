package models

type UserInfo struct {
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	UserBirthday string `json:"user_birthday,omitempty"`
	UserEmail    string `json:"user_email,omitempty"`
	UserPhone    string `json:"user_phone,omitempty"`
}
