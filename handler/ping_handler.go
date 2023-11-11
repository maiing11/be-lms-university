package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Mantap")
}
