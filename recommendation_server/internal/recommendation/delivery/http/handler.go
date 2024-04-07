package http

import (
    "encoding/json"
    "net/http"
    "recommendation_server/internal/recommendation/usecase"
)

func RecommendationHandler(uc usecase.RecommendationUsecase) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Parse user ID from query parameters or body
        userID := r.URL.Query().Get("user_id")

        // Call the usecase to get recommendations
        recommendations, err := uc.GenerateRecommendations(userID)
        if err != nil {	
            // Handle error
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Send the recommendations back as JSON
        json.NewEncoder(w).Encode(recommendations)
    }
}
