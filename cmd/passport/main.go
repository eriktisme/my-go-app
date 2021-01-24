package main

import (
	"github.com/eriktisme/passport/internal/user/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	SetupServer().Run()
}

func SetupServer() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", handlers.Ping)

	router.POST("/login", handlers.Login)

	return router
}