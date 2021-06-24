package entity

import "gorm.io/gorm"

type WebsitePassword struct {
	gorm.Model
	UserID  int    `json:"user_id"`
	Website string `json:"website"`
	Title   string `json:"title"`
}

type UpdateWebsitePassword struct {
	Website string `json:"website"`
	Title   string `json:"title"`
}

type CreateWebsitePassword struct {
	UserID  int    `json:"user_id"`
	Website string `json:"website"`
	Title   string `json:"title"`
}
