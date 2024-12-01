package models

import "time"

type User struct {
	ID        string    `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password_hash"`
	Grade     string    `json:"grade" db:"grade"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
