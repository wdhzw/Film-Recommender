package middleware

import (
	"log"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

func PanicMiddleware(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(c, "Panic occurred: %v", r)
			buf := make([]byte, 4096)
			n := runtime.Stack(buf, false)
			log.Println(c, "Stack trace: %s", string(buf[:n]))
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}()
	c.Next()
}
