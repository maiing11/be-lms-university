package main

import (
	"enigmacamp.com/be-lms-university/handler"
	"enigmacamp.com/be-lms-university/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//	intercept semuanya / routernya
	r.Use(middleware.SimpleMiddleware())

	// test ping
	r.GET("/ping", handler.PingHandler)

	// intercept satu route
	r.GET("/ping", middleware.SimpleMiddleware(), handler.PingHandler)

	r.POST("/users", handler.CreateUCHandler)
	r.POST("/users/register", handler.CreateUCWithPhotoHandler)

	r.Run(":8080")
}
