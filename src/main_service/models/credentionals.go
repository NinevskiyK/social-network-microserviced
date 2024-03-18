package models

type Credentionals struct {
	UserName string `json:"user_name,omitempty"`

	// hashedPassword (SHA3-512)
	UserPassword string `json:"user_password,omitempty"`
}
