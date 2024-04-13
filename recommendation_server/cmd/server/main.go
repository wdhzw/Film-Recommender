package main

import (
    "net/http"
    "recommendation_server/config"
    myhttp "recommendation_server/internal/recommendation/delivery/http"
    "recommendation_server/internal/recommendation/infra"
    "recommendation_server/internal/recommendation/usecase"
)

func main() {
    cfg := config.GetConfig()

    userApiClient := infra.NewUserApiClient(cfg.UserAPIURL)
    movieApiClient := infra.NewMovieApiClient(cfg.MovieAPIURL)
    recommendationUsecase := usecase.NewRecommendationUsecase(userApiClient, movieApiClient)

    http.HandleFunc("/recommendations", myhttp.RecommendationHandler(recommendationUsecase))

    http.ListenAndServe(":8080", nil)
}