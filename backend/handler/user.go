package handler

import (
	"assessment/auth"
	"assessment/entity"
	"assessment/helper"
	"assessment/layer/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

//SHOW ALL USER
func (h *userHandler) ShowAllUserHandler(c *gin.Context) {
	user, err := h.userService.SShowAllUser()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"errors": err.Error()})

		c.JSON(500, responseError)
		return
	}

	respose := helper.APIResponse("success get all user", 200, "status OK", user)
	c.JSON(200, respose)
}

// REGISTER USER
func (h *userHandler) RegisterUserHandler(c *gin.Context) {
	var inputUser entity.UserInput

	if err := c.ShouldBindJSON(&inputUser); err != nil {

		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}
}
