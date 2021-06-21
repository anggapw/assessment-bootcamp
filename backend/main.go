package main

import (
	"assessment/handler"
	"assessment/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(handler.CORSMiddleware())

	routes.UserRoute(r)

	r.Run()
}
