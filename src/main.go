package main

import (
	"github.com/gin-gonic/gin"
	"passport/src/handlers"
)

func main() {
	setupServer().Run()
}

func setupServer() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", handlers.Ping)

	router.POST("/login", handlers.Login)

	return router
}
