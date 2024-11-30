package models

type User struct {
	ID       string `json:"id" db:"id_uuid"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password,omitempty" db:"password_hash"`
	Grade    string `json:"grade" db:"grade"`
}
