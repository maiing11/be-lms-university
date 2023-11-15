package main

import (
	"enigmacamp.com/be-lms-university/handler"
	"enigmacamp.com/be-lms-university/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//	intercept semuanya / routernya
	// r.Use(middleware.SimpleMiddleware())

	v1 := r.Group("/api/v1")

	userGroup := v1.Group("/users")

	userGroup.Use(middleware.SimpleMiddleware())

	// test ping
	v1.GET("/ping", handler.PingHandler)

	// intercept satu route
	// r.GET("/ping", middleware.SimpleMiddleware(), handler.PingHandler)

	userGroup.POST("/", handler.CreateUCHandler)
	userGroup.POST("/register", handler.CreateUCWithPhotoHandler)

	r.Run(":8080")
}
