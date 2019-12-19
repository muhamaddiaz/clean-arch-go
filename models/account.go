package models

import "github.com/jinzhu/gorm"

type Account struct {
	gorm.Model
	FullName string `json:"full_name"`
	UserName string `json:"user_name"`
	Email string `json:"email"`
	Password string `json:"password"`
}
