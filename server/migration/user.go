package migration

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email            string `gorm:"unique"`
	Password         string
	FullName         string
	Address          string
	WebsitePasswords []WebsitePassword
}
