package middleware

import "github.com/gin-gonic/gin"

// TODO
/*
1. Untuk membuat sebuah middleware, kita harus membuat sebuah function yang mengembalikan sebuah handlerFunc, handlerFunc sudah disediakan oleh gin
2. Kita siapkan sebuah proses yang namanya before request & after request
3. Dijembatani oleh keyword ctx.Next() (ini wajib ada)
4. Setelah middlewarenya dibuat, kita bisa panggil .Use(middleware)
*/

func SimpleMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.WriteString("Before request\n")

		ctx.Next()

		ctx.Writer.WriteString("After request\n")
	}
}
