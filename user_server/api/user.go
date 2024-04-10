package api

import (
	"fmt"
	"net/http"

	"ESRS/user_server/client"
	"github.com/gin-gonic/gin"
)

type UserSynUpRequest struct {
	Email    string `json:"email" binding:"required"`
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ConfirmUserSynUpRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Code     string `json:"code" binding:"required"`
}

type UserLogInRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func UserSignUp(c *gin.Context) {
	var param UserSynUpRequest
	if err := c.BindJSON(&param); err != nil {
		returnError(c, http.StatusBadRequest, "invalid request param", err, nil)
		return
	}
	cognitoClient := client.NewCognitoClient("us-east-1", "3nulk6qgc8vnm668t7s0lel1u1")
	//err, result := cognitoClient.SignUp("testAccount1", "licr54@hotmail.com", "TestPassword123!")
	err, result := cognitoClient.SignUp(param.UserName, param.Email, param.Password)
	if err != nil {
		returnError(c, http.StatusInternalServerError, "sign up failed", err, nil)
		return
	}
	fmt.Println("result is", result)
	returnOK(c, result)
}

func ConfirmUserSignUp(c *gin.Context) {
	var param ConfirmUserSynUpRequest
	if err := c.BindJSON(&param); err != nil {
		returnError(c, http.StatusBadRequest, "invalid request param", err, nil)
		return
	}
	cognitoClient := client.NewCognitoClient("us-east-1", "3nulk6qgc8vnm668t7s0lel1u1")
	//err, result := cognitoClient.ConfirmSignUp("testAccount1", "999236")
	err, result := cognitoClient.ConfirmSignUp(param.UserName, param.Code)
	if err != nil {
		returnError(c, http.StatusInternalServerError, "confirm sign up failed", err, nil)
		return
	}
	fmt.Println("result is", result)
	returnOK(c, result)
}

func UserLogIn(c *gin.Context) {
	var param UserLogInRequest
	if err := c.BindJSON(&param); err != nil {
		returnError(c, http.StatusBadRequest, "invalid request param", err, nil)
		return
	}
	cognitoClient := client.NewCognitoClient("us-east-1", "3nulk6qgc8vnm668t7s0lel1u1")
	//err, result, initiateAuthOutput := cognitoClient.LogIn("testAccount1", "TestPassword123!")
	err, result, initiateAuthOutput := cognitoClient.LogIn(param.UserName, param.Password)
	if err != nil {
		returnError(c, http.StatusInternalServerError, "Log in failed", err, nil)
		return
	}
	fmt.Println("result is", result, " token: ", *initiateAuthOutput.AuthenticationResult.IdToken)
	returnOK(c, *initiateAuthOutput.AuthenticationResult.IdToken)
}
