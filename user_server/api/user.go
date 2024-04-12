package api

import (
	"fmt"
	"net/http"

	"ESRS/user_server/client"
	"ESRS/user_server/dao"
	"ESRS/user_server/utils/encoding"
	"github.com/gin-gonic/gin"
)

type UserSynUpRequest struct {
	Email    string `json:"email" binding:"required"`
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ConfirmUserSynUpRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Code     string `json:"code" binding:"required"`
}

type UserLogInRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func UserSignUp(c *gin.Context) {
	var param UserSynUpRequest
	if err := c.BindJSON(&param); err != nil {
		returnError(c, http.StatusBadRequest, "invalid request param", err, nil)
		return
	}
	cognitoClient := client.GetCognitoClient()
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
	cognitoClient := client.GetCognitoClient()
	//err, result := cognitoClient.ConfirmSignUp("testAccount1", "999236")
	err, result := cognitoClient.ConfirmSignUp(param.UserName, param.Code)
	if err != nil {
		returnError(c, http.StatusInternalServerError, "confirm sign up failed", err, nil)
		return
	}
	fmt.Println("cognito confirm sign up result is", result)
	userDao := dao.GetUserTableDAO()
	uid, err := userDao.Create(c, param.UserName, param.Email)
	if err != nil {
		returnError(c, http.StatusInternalServerError, "create user failed", err, nil)
		return
	}
	fmt.Println("user created, uid is", uid)
	returnOK(c, uid)
}

func UserLogIn(c *gin.Context) {
	var param UserLogInRequest
	if err := c.BindJSON(&param); err != nil {
		returnError(c, http.StatusBadRequest, "invalid request param", err, nil)
		return
	}
	cognitoClient := client.GetCognitoClient()
	//err, result, initiateAuthOutput := cognitoClient.LogIn("testAccount1", "TestPassword123!")
	err, result, initiateAuthOutput := cognitoClient.LogIn(param.UserName, param.Password)
	if err != nil {
		fmt.Printf("[UserLogIn] cognitoClient LogIn failed, err = %v", err)
		returnError(c, http.StatusInternalServerError, "Log in failed", err, nil)
		return
	}
	fmt.Println("result is", result, " token: ", *initiateAuthOutput.AuthenticationResult.IdToken)
	userDao := dao.GetUserTableDAO()
	user, err := userDao.GetByEmail(c, param.Email)
	if err != nil {
		returnError(c, http.StatusInternalServerError, "get user failed", err, nil)
		return
	}
	fmt.Println("user is", encoding.EncodeIgnoreError(user))
	returnOK(c, *initiateAuthOutput.AuthenticationResult.IdToken)
}
