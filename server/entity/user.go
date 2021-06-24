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
	FullName string `json:"full_name"`
	Address  string `json:"address"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCreate struct {
	FullName string `json:"full_name" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
