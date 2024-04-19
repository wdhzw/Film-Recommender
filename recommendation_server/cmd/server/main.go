package main

import (
	"CS5224_ESRS/recommendation_server/config"
	"CS5224_ESRS/recommendation_server/internal/recommendation/delivery/http"
	"CS5224_ESRS/recommendation_server/internal/recommendation/infra"
	"CS5224_ESRS/recommendation_server/internal/recommendation/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.GetConfig()

	// Initialize infrastructure
	userApiClient := infra.NewUserApiClient(cfg.UserAPIURL)
	movieApiClient := infra.NewMovieApiClient(cfg.MovieAPIURL)

	// Initialize usecase
	recommendationUsecase := usecase.NewRecommendationUsecase(userApiClient, movieApiClient)

	// Create Gin router
	r := gin.Default()

	// Apply CORS middleware
	r.Use(gin.Recovery(), http.CORSMiddleware())

	// Define routes
	r.GET("/recommendations", http.RecommendationHandler(recommendationUsecase))

	// Start the server
	r.Run(":8080")
}
