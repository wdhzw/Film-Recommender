package main

import (
    "log"
    "net/http"
    recommendationHttp "recommendation_server/internal/recommendation/delivery/http"
    "recommendation_server/internal/recommendation/usecase"
    "recommendation_server/internal/recommendation/infra"
)

func main() {
    // Initialize the recommendation use case
    recommendationUsecase := usecase.NewRecommendationUsecase(infra.UserApiClient{}, infra.MovieApiClient{})

    // Set up the HTTP server and define the route
    router := http.NewServeMux()

    // Using the alias defined in the import to refer to our custom http package
    router.HandleFunc("/recommendations", recommendationHttp.RecommendationHandler(recommendationUsecase))

    // Start the server
    log.Println("Server is running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
