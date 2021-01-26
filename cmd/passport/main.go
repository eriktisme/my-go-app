package main

import (
	handler "github.com/eriktisme/passport/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	SetupServer().Run()
}

func SetupServer() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", handler.Ping)


	router.POST("/login", handler.Login)

	return router
}