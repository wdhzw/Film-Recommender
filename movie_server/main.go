package movie_server

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// init infra
	InitMysql()

	r := gin.Default()

	r.Use(gin.Recovery())

	r.GET("/movie_server//:movie_id", GetMovieDetailsById)
	r.GET("/movie_server/popular", GetPopularMovies)
	r.GET("/movie_server/search", Search)
	r.POST("/movie_server/update", UpdateMovieMeta)

	r.Run() // listen and serve on 0.0.0.0:8080
}
