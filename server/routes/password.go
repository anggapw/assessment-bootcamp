package routes

import (
	"password-app/handler"
	"password-app/layer/password"

	"github.com/gin-gonic/gin"
)

var (
	passwordRepository = password.NewRepository(DB)
	passwordService    = password.NewService(passwordRepository)
	passwordHandler    = handler.NewPasswordhandler(passwordService, authService)
)

func PasswordRoute(r *gin.Engine) {
	r.POST("/password", handler.Middleware(userService, authService), passwordHandler.CreatePasswordHandler)
	r.GET("/password", handler.Middleware(userService, authService), passwordHandler.GetAllPasswordhandler)
	r.GET("/password/:id", handler.Middleware(userService, authService), passwordHandler.ShowPasswordByIDhandler)
	r.PUT("/password/:id", handler.Middleware(userService, authService), passwordHandler.UpdatePasswordByIDHandler)
	r.DELETE("/password/:id", handler.Middleware(userService, authService), passwordHandler.DeleteByPasswordIDHandler)
}
