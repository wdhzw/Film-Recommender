package usecase

import (
	"fmt"
	"sort"

	"recommendation_server/internal/recommendation/infra"
)

type RecommendationUsecase struct {
	userApiClient  *infra.UserApiClient
	movieApiClient *infra.MovieApiClient
}

func NewRecommendationUsecase(userApiClient *infra.UserApiClient, movieApiClient *infra.MovieApiClient) *RecommendationUsecase {
	return &RecommendationUsecase{
		userApiClient:  userApiClient,
		movieApiClient: movieApiClient,
	}
}

type Item struct {
	Movie interface{}
	Score float64
}

func (uc *RecommendationUsecase) GeneratePersonalizedRecommendations(email string, pageNumber int) ([]interface{}, error) {
	user, err := uc.userApiClient.FetchUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("error fetching user data: %v", err)
	}
    if user == nil || len(user.PreferredGenre) == 0 {
        // Check if the user object is valid or if preferred genres are empty
        return nil, fmt.Errorf("no user found with the email %s or no preferred genres set", email)
    }

    popularResponse, err := uc.movieApiClient.GetPopularMovies(pageNumber)
    if err != nil {
        return nil, fmt.Errorf("error fetching popular movies: %v", err)
    }
    highRateResponse, err := uc.movieApiClient.GetHighRateMovies(pageNumber)
    if err != nil {
        return nil, fmt.Errorf("error fetching high-rate movies: %v", err)
    }

	preferenceSet := make(map[string]struct{})
	for _, genre := range user.PreferredGenre {
		preferenceSet[genre] = struct{}{}
	}

	movies := make(map[int64]*Item)
	for _, movie := range popularResponse.Content {
		mid := movie.MovieID
		if _, exists := movies[mid]; !exists {
			movies[mid] = &Item{
				Movie: movie,
				Score: movie.Popularity,
			}
		}

		for _, genre := range movie.Genres {
			if _, ok := preferenceSet[genre]; ok {
				movies[mid].Score += 5
			}
		}
	}

	for _, movie := range highRateResponse.Content {
		mid := movie.MovieID
		if item, exists := movies[mid]; exists {
			item.Score += movie.Rate * 2
		} else {
			movies[mid] = &Item{
				Movie: movie,
				Score: movie.Rate * 2,
			}
		}

		for _, genre := range movie.Genres {
			if _, ok := preferenceSet[genre]; ok {
				movies[mid].Score += 5
			}
		}
	}

	rawItems := make([]*Item, 0, len(movies))
	for _, item := range movies {
		rawItems = append(rawItems, item)
	}

	sort.Slice(rawItems, func(i, j int) bool {
		return rawItems[i].Score > rawItems[j].Score // Sort in descending order
	})

	results := make([]interface{}, 0, len(rawItems))
	for _, item := range rawItems {
		results = append(results, item.Movie)
	}

	return results, nil
}
