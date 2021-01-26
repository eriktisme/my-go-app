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

	router.GET("/csrf-cookie", handler.Csrf)

	router.POST("/login", handler.Login)

	return router
}