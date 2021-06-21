package routes

import (
	"assessment/auth"
	"assessment/config"
	"assessment/handler"
	"assessment/layer/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB = config.Connection()
	authService             = auth.NewService()
	userService             = user.NewService(userRepository)
	userRepository          = user.NewRepository(DB)
	userHandler             = handler.NewUserHandler(userService, authService)
)

func UserRoute(r *gin.Engine) {
	r.GET("/users", handler.Middleware(userService, authService), userHandler.ShowAllUserHandler)
	r.POST("/users/register", userHandler.RegisterUserHandler)
	r.POST("/users/login", userHandler.LoginUserHandler)
}
