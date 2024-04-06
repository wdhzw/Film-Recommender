package service

import (
	"CS5224_ESRS/movie/model"
	"CS5224_ESRS/movie/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type Actor struct {
	Name       string `json:"name" form:"movie_id"`
	Character  string `json:"character" form:"movie_id"`
	ProfileUri string `json:"profile_uri" form:"movie_id"`
}

type MovieDetails struct {
	MovieId             int64     `json:"movie_id"`
	Title               string    `json:"title"`
	Overview            string    `json:"overview"`
	Rate                float64   `json:"rate"`
	Popularity          float64   `json:"popularity"`
	Homepage            string    `json:"homepage"`
	PosterURL           string    `json:"poster_uri"`
	Actors              []*Actor  `json:"actors"`
	Director            string    `json:"director"`
	Writers             []string  `json:"writers"`
	Genres              []string  `json:"genres"`
	ProductionCountries []string  `json:"production_countries"`
	Languages           []string  `json:"language"`
	ReleaseDate         time.Time `json:"release_date"`
	Duration            int       `json:"duration"`
	Keywords            []string  `json:"keywords"`
}

func GetMovieDetailsById(c *gin.Context) {

	movieIdStr, isExist := c.Params.Get("movie_id")

	if !isExist || movieIdStr == "" {
		returnError(c, 200, ErrorMovieIdInvalid, "MovieId is empty, please check your system")
		return
	}

	movieId, _ := strconv.ParseInt(movieIdStr, 10, 64)
	movie, err := MovieModel.GetMovieModelById(movieId)
	if err != nil {
		returnError(c, 500, InternalError, "DB has some unexpected errors, please contact with developers to check")
		return
	}

	resp := packMovieDetails(movie)

	// pack response
	returnOK(c, resp)
	return

}

func packMovieDetails(model *model.MovieModel) *MovieDetails {
	resp := &MovieDetails{
		MovieId:             model.MovieID,
		Title:               model.Title,
		Overview:            model.Overview,
		Rate:                model.Rate,
		Popularity:          model.Popularity,
		Homepage:            model.Homepage,
		PosterURL:           utils.FillURI(model.PosterURI),
		Director:            model.Director,
		Writers:             utils.SplitSeries(model.Writers),
		Genres:              utils.SplitSeries(model.Genres),
		ProductionCountries: utils.SplitSeries(model.ProductionCountry),
		Languages:           utils.SplitSeries(model.Language),
		ReleaseDate:         model.ReleaseDate,
		Duration:            model.Duration,
		Keywords:            utils.SplitSeries(model.Keyword),
	}

	actors := make([]*Actor, 0)
	_ = json.Unmarshal(model.Actors, &actors)
	resp.Actors = actors

	for _, actor := range resp.Actors {
		if actor.ProfileUri == "" {
			continue
		}
		actor.ProfileUri = utils.FillURI(actor.ProfileUri)
	}

	return resp
}
