package main

import (
	"password-app/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(handler.CORSMiddleware())

	r.POST("/user", handler.HandlePostUser)
	r.GET("/users", handler.HandleGetUser)
	r.GET("/users/:id", handler.HandleGetUserByID)
	r.PUT("/user/:id", handler.HandleUpdateUser)
	r.DELETE("/user/:id", handler.HandleDeleteUser)

	r.POST("/password", handler.HandlePostPassword)
	r.GET("/passwords", handler.HandleGetPassword)
	r.GET("/password/:id", handler.HandleGetPasswordByID)
	r.PUT("/password/:id", handler.HandleUpdatePassword)
	r.DELETE("/password/:id", handler.HandleDeletePassword)

	r.Run()

}
