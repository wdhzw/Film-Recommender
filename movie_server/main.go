package main

import (
	"CS5224_ESRS/movie/infra"
	"CS5224_ESRS/movie/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// init infra
	infra.InitMysql()
	service.InitService()

	r := gin.Default()

	r.Use(gin.Recovery())

	r.GET("/movie_server/:movie_id", service.GetMovieDetailsById)
	r.GET("/movie_server/popular", service.GetPopularMovies)
	r.GET("/movie_server/top_rate", service.GetHighRateMovies)
	r.GET("/movie_server/search", service.Search)
	r.POST("/movie_server/update", service.UpdateMovieMeta)

	r.Run() // listen and serve on 0.0.0.0:8080
}
