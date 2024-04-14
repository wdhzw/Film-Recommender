package usecase

import (
    "fmt"
    "strings"

    "recommendation_server/internal/recommendation/infra"
)

type RecommendationUsecase struct {
    userApiClient   *infra.UserApiClient
    movieApiClient  *infra.MovieApiClient
}

func NewRecommendationUsecase(userApiClient *infra.UserApiClient, movieApiClient *infra.MovieApiClient) *RecommendationUsecase {
    return &RecommendationUsecase{
        userApiClient:   userApiClient,
        movieApiClient:  movieApiClient,
    }
}

func (uc *RecommendationUsecase) GeneratePersonalizedRecommendations(email string) ([]infra.MovieModel, error) {
    // Fetch user data
    user, err := uc.userApiClient.FetchUserByEmail(email)
    if err != nil {
        return nil, fmt.Errorf("error fetching user data: %v", err)
    }

    // Fetch popular and high-rated movies
    popularMovies, err := uc.movieApiClient.GetPopularMovies()
    if err != nil {
        return nil, fmt.Errorf("error fetching popular movies: %v", err)
    }
    highRateMovies, err := uc.movieApiClient.GetHighRateMovies()
    if err != nil {
        return nil, fmt.Errorf("error fetching high-rate movies: %v", err)
    }

    // Combine popular and high-rated movies
    allMovies := append(popularMovies, highRateMovies...)

    // Filter movies based on user's preferred genres
    var recommendedMovies []infra.MovieModel
    for _, movie := range allMovies {
        movieGenres := strings.Split(movie.Genres, ",")
        for _, genre := range user.PreferredGenre {
            if containsGenre(movieGenres, genre) {
                recommendedMovies = append(recommendedMovies, movie)
                break
            }
        }
    }

    return recommendedMovies, nil
}

func containsGenre(genres []string, target string) bool {
    for _, genre := range genres {
        if strings.TrimSpace(genre) == target {
            return true
        }
    }
    return false
}