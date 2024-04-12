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

type UpdateUserGenreRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Genre    string `json:"genre" binding:"required"`
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

func UpdateUserGenre(c *gin.Context) {
	var param UpdateUserGenreRequest
	if err := c.BindJSON(&param); err != nil {
		returnError(c, http.StatusBadRequest, "invalid request param", err, nil)
		return
	}

	// Retrieve the current user's genres
	userDao := dao.GetUserTableDAO()
	user, err := userDao.GetByEmail(c, param.Email)
	if err != nil {
		returnError(c, http.StatusInternalServerError, "failed to retrieve user", err, nil)
		return
	}

	// Implement LRU logic for genres
	maxGenres := 5
	updatedGenres := updateUserGenres(user.PreferredGenre, param.Genre, maxGenres)

	// Update the user's genres
	updateParams := dao.UpdateUserParams{
		UID:            user.UserID,
		PreferredGenre: updatedGenres,
	}
	if err := userDao.Update(c, updateParams); err != nil {
		returnError(c, http.StatusInternalServerError, "failed to update user genres", err, nil)
		return
	}

	returnOK(c, "")
}

// updateUserGenres manages the LRU logic for the genre list
func updateUserGenres(genres []string, newGenre string, maxGenres int) []string {
	// Remove the genre if it already exists to treat it as the most recently used
	index := -1
	for i, genre := range genres {
		if genre == newGenre {
			index = i
			break
		}
	}
	if index != -1 {
		genres = append(genres[:index], genres[index+1:]...)
	}
	genres = append(genres, newGenre) // Add the new genre at the end to mark it as the most recent

	// If the list exceeds the maximum size, remove the least recently used genre (the first element)
	if len(genres) > maxGenres {
		genres = genres[1:] // Remove the first element
	}

	return genres
}
