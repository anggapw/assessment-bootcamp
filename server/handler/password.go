package handler

import (
	"net/http"
	"password-app/auth"
	"password-app/entity"
	"password-app/layer/password"

	"github.com/gin-gonic/gin"
)

type passwordHandler struct {
	passwordService password.PasswordService
	authService     auth.Service
}

func NewPasswordhandler(passwordService password.PasswordService, authService auth.Service) *passwordHandler {
	return &passwordHandler{passwordService, authService}
}

func (h *passwordHandler) CreatePasswordHandler(c *gin.Context) {
	var inputPassword entity.CreateWebsitePassword

	if err := c.ShouldBindJSON(&inputPassword); err != nil {
		c.JSON(400, gin.H{
			"error": "input data required",
		})
		return
	}

	userID := int(c.MustGet("currentUser").(int))

	response, err := h.passwordService.CreateNewPassword(userID, inputPassword)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, response)
}

func (h *passwordHandler) GetAllPasswordhandler(c *gin.Context) {
	password, err := h.passwordService.GetAllPassword()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, password)

}

func (h *passwordHandler) ShowPasswordByIDhandler(c *gin.Context) {
	id := c.Param("id")

	password, err := h.passwordService.GetPasswordByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, password)
}

func (h *passwordHandler) UpdatePasswordByIDHandler(c *gin.Context) {
	id := c.Param("id")

	var updatePasswordInput entity.CreateWebsitePassword

	if err := c.ShouldBindJSON(&updatePasswordInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	password, err := h.passwordService.UpdatePasswordByID(id, updatePasswordInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	idParam := int(password.UserID)

	userData := int(c.MustGet("currentUser").(int))

	if idParam != userData {
		c.JSON(401, gin.H{
			"error": "unauthorize user",
		})
		return
	}

	c.JSON(http.StatusOK, password)
}

func (h *passwordHandler) DeleteByPasswordIDHandler(c *gin.Context) {
	id := c.Param("id")

	password, _ := h.passwordService.GetPasswordByID(id)

	idParam := int(password.UserID)

	userData := int(c.MustGet("currentUser").(int))

	if idParam != userData {
		c.JSON(401, gin.H{
			"error": "unauthorize user",
		})
		return
	}

	book, err := h.passwordService.DeletePasswordByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, book)
}
