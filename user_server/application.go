package main

import (
	"os"

	"ESRS/user_server/api"
	"ESRS/user_server/client"
	"ESRS/user_server/config"
	"ESRS/user_server/middleware"
	"github.com/gin-gonic/gin"
)

func initDependencies() {
	config.Init()
	client.InitCognitoClient()
	client.InitDynamoDB()
}

func main() {
	os.Setenv("PORT", "8080")
	os.Setenv("GIN_MODE", "release")

	r := gin.Default()

	initDependencies()

	r.Use(
		middleware.PanicMiddleware,
	)

	public := r.Group("/api")
	{
		public.POST("/user_sign_up", api.UserSignUp)
		public.POST("/confirm_user_sign_up", api.ConfirmUserSignUp)
		public.POST("/user_login", api.UserLogIn)
		public.POST("/update_user_genre", api.UpdateUserGenre)
	}

	r.Run(":" + os.Getenv("PORT"))
}
