package http

import (
	"encoding/json"
	"net/http"

	"CS5224_ESRS/recommendation_server/internal/recommendation/usecase"
)

func RecommendationHandler(uc *usecase.RecommendationUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract user email from request body
		var req struct {
			Email string `json:"email"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		// Generate personalized recommendations
		recommendedMovies, err := uc.GeneratePersonalizedRecommendations(req.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Return the recommended movies as JSON response
		json.NewEncoder(w).Encode(recommendedMovies)
	}
}
