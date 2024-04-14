package service

import (
	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	searchWord := c.Query("search_word")

	movies, err := MovieModel.SearchMovie(searchWord)
	if err != nil {
		returnError(c, 500, InternalError, "DB has some unexpected errors, please contact with developers to check")
		return
	}

	items := make([]*PopularMovieItem, 0)
	for _, movie := range movies {
		item := packPopularMovieItem(movie)
		items = append(items, item)
	}

	c.JSON(200, gin.H{
		"content": items,
	})
	return
}
