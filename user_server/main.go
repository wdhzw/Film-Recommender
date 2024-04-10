package main

import (
	"os"

	"ESRS/user_server/api"
	"ESRS/user_server/middleware"
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
		public.POST("/user_sign_up", api.UserSignUp)
		public.POST("/confirm_user_sign_up", api.ConfirmUserSignUp)
		public.POST("/user_login", api.UserLogIn)
	}
}
