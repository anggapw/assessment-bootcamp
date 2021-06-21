package entity

import "time"

type PasswordList struct {
	ID          int       `gorm:"PrimaryKey" json:"id"`
	Website     string    `json:"website"`
	Passwordweb string    `json:"password_web"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      int       `json:"user_id"`
}

type PasswordListInput struct {
	Website     string `json:"website" binding:"required"`
	Passwordweb string `json:"password_web" binding:"required"`
}
