package handler

import (
	"log"
	"net/http"
	"password-app/entity"

	"github.com/gin-gonic/gin"
)

func HandlePostPassword(c *gin.Context) {
	var password entity.WebsitePassword

	if err := c.ShouldBindJSON(&password); err != nil {
		log.Println(err.Error())
		return
	}

	if err := DB.Create(&password).Error; err != nil {
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, password)

}

func HandleGetPassword(c *gin.Context) {
	var passwords []entity.WebsitePassword

	if err := DB.Find(&passwords).Error; err != nil {
		log.Println(err.Error())
	}

	c.JSON(http.StatusOK, passwords)
}

func HandleGetPasswordByID(c *gin.Context) {
	var password entity.WebsitePassword

	id := c.Params.ByName("id")

	if err := DB.Where("id = ?", id).Find(&password).Error; err != nil {
		log.Println(err.Error())
	}

	c.JSON(http.StatusOK, password)
}

func HandleDeletePassword(c *gin.Context) {
	id := c.Params.ByName("id")

	if err := DB.Where("id ' ?", id).Delete(&entity.WebsitePassword{}).Error; err != nil {
		log.Println(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"message": "delete password succesfully",
	})

}

func HandleUpdatePassword(c *gin.Context) {
	var password entity.WebsitePassword
	var passwordUpdate entity.UpdateWebsitePassword

	id := c.Params.ByName("id")
	DB.Where("id = ?", id).Find(&password)

	if err := c.ShouldBindJSON(&passwordUpdate); err != nil {
		log.Println(err.Error())
		return
	}

	password.Website = passwordUpdate.Website
	password.Title = passwordUpdate.Title

	if err := DB.Where("id = ?", id).Save(&password).Error; err != nil {
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, password)
}
