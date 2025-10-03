package models

import "time"

type User struct {
	ID              int64     `json:"id" db:"id"`
	Email           string    `json:"email" db:"email"`
	PasswordHash    string    `json:"-" db:"password_hash"`
	ImageURI        string    `json:"image_uri" db:"image_uri"`
	CompanyName     string    `json:"company_name" db:"company_name"`
	CompanyImageURI string    `json:"company_image_uri" db:"company_image_uri"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}
