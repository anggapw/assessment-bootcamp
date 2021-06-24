package main

import (
	"password-app/handler"
	"password-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(handler.CORSMiddleware())

	routes.UserRoute(r)
	routes.PasswordRoute(r)

	r.Run()
}
