package handler

import (
	"log"
	"net/http"
	"password-app/config"
	"password-app/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB = config.Connect()

func HandlePostUser(c *gin.Context) {
	var user entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err.Error())
		return
	}

	if err := DB.Create(&user).Error; err != nil {
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, user)

}

func HandleGetUser(c *gin.Context) {
	var users []entity.User

	if err := DB.Find(&users).Error; err != nil {
		log.Println(err.Error())
	}

	c.JSON(http.StatusOK, users)
}

func HandleGetUserByID(c *gin.Context) {
	var user entity.User

	id := c.Params.ByName("id")

	if err := DB.Where("id = ?", id).Find(&user).Error; err != nil {
		log.Println(err.Error())
	}

	c.JSON(http.StatusOK, user)
}

func HandleDeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")

	if err := DB.Where("id ' ?", id).Delete(&entity.User{}).Error; err != nil {
		log.Println(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"message": "delete password succesfully",
	})

}

func HandleUpdateUser(c *gin.Context) {
	var user entity.User
	var userUpdate entity.UserUpdate

	id := c.Params.ByName("id")
	DB.Where("id = ?", id).Find(&user)

	if err := c.ShouldBindJSON(&userUpdate); err != nil {
		log.Println(err.Error())
		return
	}

	user.Email = userUpdate.Email
	user.FullName = userUpdate.FullName
	user.Address = userUpdate.Address

	if err := DB.Where("id = ?", id).Save(&user).Error; err != nil {
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
