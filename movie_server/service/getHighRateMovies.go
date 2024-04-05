package service

import (
	"CS5224_ESRS/movie/model"
	"CS5224_ESRS/movie/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type TopRateMovieItem struct {
	MovieId     int64     `json:"movie_id"`
	Title       string    `json:"title"`
	Rate        float64   `json:"rate"`
	Genres      []string  `json:"genres"`
	ReleaseDate time.Time `json:"release_date"`
	Duration    int       `json:"duration"`
	Keywords    []string  `json:"keywords"`
}

func packTopRateMovieItem(model *model.MovieModel) *TopRateMovieItem {
	return &TopRateMovieItem{
		MovieId:     model.MovieID,
		Title:       model.Title,
		Rate:        model.Rate,
		Genres:      utils.SplitSeries(model.Genres),
		ReleaseDate: model.ReleaseDate,
		Duration:    model.Duration,
		Keywords:    utils.SplitSeries(model.Keyword),
	}
}

func GetHighRateMovies(c *gin.Context) {
	pageStr := c.Query("page")
	if pageStr == "" {
		pageStr = "1"
	}

	page, _ := strconv.Atoi(pageStr)

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

	totalPages := total / int64(10)
	if total/int64(10) != 0 {
		totalPages++
	}
	items := make([]*TopRateMovieItem, 0)
	for _, movie := range movies {
		item := packTopRateMovieItem(movie)
		items = append(items, item)
	}

	c.JSON(200, gin.H{
		"content":      items,
		"total_pages":  totalPages,
		"total_number": total,
		"current_page": page,
	})
	return
}
