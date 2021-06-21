package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
}

type UserUpdate struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
