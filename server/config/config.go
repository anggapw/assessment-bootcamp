package config

import (
	"fmt"
	"log"
	"os"
	"password-app/migration"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DBUSER := os.Getenv("DB_USER")
	DBPASSWORD := os.Getenv("DB_PASSWORD")
	DBHOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUSER, DBPASSWORD, DBHOST, DBNAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&migration.User{}, &migration.WebsitePassword{})
	return db
}
