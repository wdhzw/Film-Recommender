package movie_server

import (
	"CS5224_ESRS/movie/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	InternalError int = 10001

	ErrorMovieIdInvalid int = 10002
)

var (
	MovieModel = model.InitMovie()
)

func GetMovieDetailsById(c *gin.Context) {

	movieIdStr := c.GetString("movie_id")

	if movieIdStr == "" {
		returnError(c, 200, ErrorMovieIdInvalid, "MovieId is empty, please check your system")
		return
	}

	movieId, _ := strconv.ParseInt(movieIdStr, 10, 64)
	movie, err := MovieModel.GetMovieModelById(movieId)
	if err != nil {
		returnError(c, 500, InternalError, "DB has some unexpected errors, please contact with developers to check")
		return
	}

	// uri -> url
	movie.PosterURI = FillURI(movie.PosterURI)
	if movie.Actors, err = FillActorProfileImage(movie.Actors); err != nil {
		returnError(c, 500, InternalError, "Jsonify occurs unexpected errors, please check the completeness of data")
		return
	}

	// pack response
	returnOK(c, movie)
	return

}

func GetPopularMovies(c *gin.Context) {
	page := c.GetInt("page")

	movies, err := MovieModel.GetPopularMovies(page)
	if err != nil {
		returnError(c, 500, InternalError, "DB has some unexpected errors, please contact with developers to check")
		return
	}

	total, err := MovieModel.CountTotalMovies()
	if err != nil {
		returnError(c, 500, InternalError, "DB has some unexpected errors, please contact with developers to check")
		return
	}

	// 计算总页数
	totalPages := total / int64(10)
	if total/int64(10) != 0 {
		totalPages++
	}

	c.JSON(200, gin.H{
		"content": movies,
		"total":   totalPages,
	})
	return
}

func GetHighRateMovies(c *gin.Context) {
	page := c.GetInt("page")

	movies, err := MovieModel.GetHighRateMovies(page)
	if err != nil {
		returnError(c, 500, InternalError, "DB has some unexpected errors, please contact with developers to check")
		return
	}

	total, err := MovieModel.CountTotalMovies()
	if err != nil {
		returnError(c, 500, InternalError, "DB has some unexpected errors, please contact with developers to check")
		return
	}

	// 计算总页数
	totalPages := total / int64(10)
	if total/int64(10) != 0 {
		totalPages++
	}

	c.JSON(200, gin.H{
		"content": movies,
		"total":   totalPages,
	})
	return
}

func Search(c *gin.Context) {

}

// UpdateMovieMeta only update movie_server rate and popularity
func UpdateMovieMeta(c *gin.Context) {
	movieIdStr := c.GetString("movie_id")

	if movieIdStr == "" {
		returnError(c, 200, ErrorMovieIdInvalid, "MovieId is empty, please check your system")
		return
	}

	movieId, _ := strconv.ParseInt(movieIdStr, 10, 64)

	rateStr := c.GetString("rate")
	popularityStr := c.GetString("popularity")
	rate, _ := strconv.ParseFloat(rateStr, 64)
	popularity, _ := strconv.ParseFloat(popularityStr, 64)

	err := MovieModel.UpdateMovies(movieId, popularity, rate)
	if err != nil {
		returnError(c, 500, InternalError, "DB has some unexpected errors, please contact with developers to check")
		return
	}

	returnOK(c, "ok")
	return
}

func returnError(c *gin.Context, status int, businessErrorCode int, message string) {
	c.JSON(status, map[string]interface{}{"error_msg": message, "status_code": businessErrorCode})
}

func returnOK(c *gin.Context, content interface{}) {
	c.JSON(200, map[string]interface{}{"error": nil, "content": content, "status_code": 0})
}
