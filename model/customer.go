package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}