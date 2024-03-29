package api

import (
	"github.com/gin-gonic/gin"
)

const (
	InternalError int = 10001

	UserNotRegister        int = 1000
	InvitationCodeNotValid int = 1001

	RetweetVerifyRateLimit int = 2000
	XAccountAlreadyInUse   int = 2001
	XAccountFailAuth       int = 2002
)

func returnError(c *gin.Context, status int, message string, err error, businessErrorCode *int) {
	code := InternalError
	if businessErrorCode != nil {
		code = *businessErrorCode
	}
	c.JSON(status, map[string]interface{}{"error_msg": message, "data": nil, "status_code": code})
}

func returnOK(c *gin.Context, resp interface{}) {
	c.JSON(200, map[string]interface{}{"error": nil, "data": resp, "status_code": 0})
}
