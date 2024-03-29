package main

import (
	"ESRS/api"
	"ESRS/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func initDependencies() {
}

func main() {
	os.Setenv("PORT", "8080")
	os.Setenv("GIN_MODE", "release")

	r := gin.Default()

	r.Use(
		middleware.PanicMiddleware,
	)

	public := r.Group("/api")
	{
		public.POST("/user_login", api.UserLogin)
	}

}
