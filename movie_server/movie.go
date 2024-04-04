package movie_server

import (
	"CS5224_ESRS/movie/model"
	"github.com/gin-gonic/gin"
)

const (
	InternalError int = 10001

	ErrorMovieIdInvalid int = 10002
)

var (
	MovieModel = model.InitMovie()
)

func GetMovieDetailsById(c *gin.Context) {

	movieId := c.GetInt64("movie_id")

	if movieId == 0 {
		returnError(c, 200, ErrorMovieIdInvalid, "MovieId is empty, please check your system")
		return
	}

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

}

func Search(c *gin.Context) {

}

// UpdateMovieMeta only update movie_server rate and popularity
func UpdateMovieMeta(c *gin.Context) {
}

func returnError(c *gin.Context, status int, businessErrorCode int, message string) {
	c.JSON(status, map[string]interface{}{"error_msg": message, "status_code": businessErrorCode})
}

func returnOK(c *gin.Context, content interface{}) {
	c.JSON(200, map[string]interface{}{"error": nil, "content": content, "status_code": 0})
}
