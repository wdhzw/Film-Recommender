package service

import (
	"CS5224_ESRS/movie/model"
	"github.com/gin-gonic/gin"
)

const (
	InternalError int = 10001

	ErrorMovieIdInvalid         int = 10002
	ErrorUpdateMovieFormInvalid int = 10003
)

var (
	MovieModel *model.MovieProxy
)

func InitService() {
	MovieModel = model.InitMovie()
}

func returnError(c *gin.Context, status int, businessErrorCode int, message string) {
	c.JSON(status, map[string]interface{}{"error_msg": message, "status_code": businessErrorCode})
}

func returnOK(c *gin.Context, content interface{}) {
	c.JSON(200, map[string]interface{}{"error": nil, "content": content, "status_code": 0})
}
