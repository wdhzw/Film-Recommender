package service

import (
	"CS5224_ESRS/movie/model"
	"CS5224_ESRS/movie/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CreateMovieRequest struct {
	MovieID             int64     `form:"movie_id"`
	Title               string    `form:"title"`
	Overview            string    `form:"overview"`
	Rate                float64   `form:"rate"`
	Popularity          float64   `form:"popularity"`
	Homepage            string    `form:"homepage"`
	PosterURI           string    `form:"poster_uri"`
	Actors              []*Actor  `form:"actors"`
	Director            string    `form:"director"`
	Writers             []string  `form:"writers"`
	Genres              []string  `form:"genres"`
	ProductionCountries []string  `form:"production_countries"`
	Languages           []string  `form:"languages"`
	ReleaseDate         time.Time `form:"release_date"`
	Duration            int       `form:"duration"`
	Keywords            []string  `form:"keywords"`
}

func CreateMovie(c *gin.Context) {
	req := &CreateMovieRequest{}
	if err := c.ShouldBind(&req); err != nil {
		returnError(c, http.StatusBadRequest, ErrorUpdateMovieFormInvalid, "Create Form invalid, please check your parameters")
	}

	mm := convertMovieToModel(req)
	err := MovieModel.CreateMovie(mm)
	if err != nil {
		returnError(c, 500, InternalError, "DB has some unexpected errors, please contact with developers to check")
		return
	}
	returnOK(c, "ok")
	return
}

func convertMovieToModel(req *CreateMovieRequest) *model.MovieModel {
	movieModel := &model.MovieModel{
		MovieID:           req.MovieID,
		Title:             req.Title,
		Overview:          req.Overview,
		Rate:              req.Rate,
		Popularity:        req.Popularity,
		Homepage:          req.Homepage,
		PosterURI:         req.PosterURI,
		Director:          req.Director,
		Writers:           utils.JointSeries(req.Writers),
		Genres:            utils.JointSeries(req.Genres),
		ProductionCountry: utils.JointSeries(req.ProductionCountries),
		Language:          utils.JointSeries(req.Languages),
		ReleaseDate:       req.ReleaseDate,
		Duration:          req.Duration,
		Keyword:           utils.JointSeries(req.Keywords),
	}

	actorBytes, _ := json.Marshal(req.Actors)
	movieModel.Actors = actorBytes

	return movieModel
}
