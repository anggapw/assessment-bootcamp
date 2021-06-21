package migration

import "time"

type User struct {
	ID           int `gorm:"PrimaryKey"`
	FullName     string
	Address      string
	Email        string `gorm:"unique"`
	Password     string
	Passwordlist []PasswordList `gorm:"ForeignKey:UserID"`
}

type PasswordList struct {
	ID        int `gorm:"PrimaryKey"`
	Website   string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int
}
