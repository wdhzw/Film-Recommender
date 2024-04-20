package http

import (
    "net/http"
    "recommendation_server/internal/recommendation/usecase"
    "github.com/gin-gonic/gin"
)

func RecommendationHandler(uc *usecase.RecommendationUsecase) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Extract user email from query parameters
        email := c.Query("email")
        if email == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
            return
        }

        // Generate personalized recommendations
        recommendedMovies, err := uc.GeneratePersonalizedRecommendations(email)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        // Return the recommended movies as JSON response
        c.JSON(http.StatusOK, recommendedMovies)
    }
}