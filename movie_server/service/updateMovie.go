package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateRequest struct {
	MovieId    string `form:"movie_id"`
	Rate       string `form:"rate"`
	Popularity string `form:"popularity"`
}

// UpdateMovieMeta only update movie_server rate and popularity
func UpdateMovieMeta(c *gin.Context) {
	req := &UpdateRequest{}
	if err := c.ShouldBind(&req); err != nil {
		returnError(c, http.StatusBadRequest, ErrorUpdateMovieFormInvalid, "Update Form invalid, please check your parameters")
	}

	movieId, _ := strconv.ParseInt(req.MovieId, 10, 64)
	rate, _ := strconv.ParseFloat(req.Rate, 64)
	popularity, _ := strconv.ParseFloat(req.Popularity, 64)

	err := MovieModel.UpdateMovies(movieId, popularity, rate)
	if err != nil {
		returnError(c, 500, InternalError, "DB has some unexpected errors, please contact with developers to check")
		return
	}

	returnOK(c, "ok")
	return
}
