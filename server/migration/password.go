package migration

import "gorm.io/gorm"

type WebsitePassword struct {
	gorm.Model
	UserID  uint
	Website string
	Title   string
}
