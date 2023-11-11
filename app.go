package main

import (
	"enigmacamp.com/be-lms-university/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// test ping
	r.GET("/ping", handler.PingHandler)

	r.POST("/users", handler.CreateUCHandler)
	r.POST("/users/register", handler.CreateUCWithPhotoHandler)

	r.Run(":8080")
}
