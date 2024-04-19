package http

import (
    "net/http"
    "strconv"
    "recommendation_server/internal/recommendation/usecase"
    "github.com/gin-gonic/gin"
)

func RecommendationHandler(uc *usecase.RecommendationUsecase) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Extract user email from query parameters
        email := c.Query("email")
        if email == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
            return
        }

        // Extract page number from query parameters
        page := c.DefaultQuery("page", "1")  // Default to page 1 if not specified
        pageNumber, err := strconv.Atoi(page)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
            return
        }

        // Generate personalized recommendations
        recommendedMovies, err := uc.GeneratePersonalizedRecommendations(email, pageNumber)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        // Return the recommended movies as JSON response
        c.JSON(http.StatusOK, recommendedMovies)
    }
}
