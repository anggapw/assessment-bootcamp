package entity

import "gorm.io/gorm"

type WebsitePassword struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	Website string `json:"website"`
	Title   string `json:"title"`
}

type UpdateWebsitePassword struct {
	Website string `json:"website"`
	Title   string `json:"title"`
}
