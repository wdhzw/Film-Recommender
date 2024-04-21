package service

import (
	"CS5224_ESRS/movie/model"
	"CS5224_ESRS/movie/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type PopularMovieItem struct {
	MovieId     int64     `json:"movie_id"`
	Title       string    `json:"title"`
	Popularity  float64   `json:"popularity"`
	Genres      []string  `json:"genres"`
	ReleaseDate time.Time `json:"release_date"`
	Duration    int       `json:"duration"`
	Keywords    []string  `json:"keywords"`
	PosterURL   string    `json:"poster_uri"`
}

func packPopularMovieItem(model *model.MovieModel) *PopularMovieItem {
	return &PopularMovieItem{
		MovieId:     model.MovieID,
		Title:       model.Title,
		Popularity:  model.Popularity,
		Genres:      utils.SplitSeries(model.Genres),
		ReleaseDate: model.ReleaseDate,
		Duration:    model.Duration,
		Keywords:    utils.SplitSeries(model.Keyword),
		PosterURL:   utils.FillURI(model.PosterURI),
	}
}

func GetPopularMovies(c *gin.Context) {
	pageStr := c.Query("page")
	if pageStr == "" {
		pageStr = "1"
	}

	page, _ := strconv.Atoi(pageStr)
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

	items := make([]*PopularMovieItem, 0)
	for _, movie := range movies {
		item := packPopularMovieItem(movie)
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
