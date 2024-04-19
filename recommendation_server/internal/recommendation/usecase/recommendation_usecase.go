package usecase

import (
	"fmt"
	"sort"
	"strings"

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
	Movie infra.MovieModel
	Score float64
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

	preferences := user.PreferredGenre
	preferenceSet := make(map[string]struct{}, 0)
	for _, p := range preferences {
		if _, ok := preferenceSet[p]; !ok {
			preferenceSet[p] = struct{}{}
		}
	}

	movies := make(map[int64]*Item, 0)
	for _, popularMovie := range popularMovies {
		mid := popularMovie.MovieID

		if _, ok := movies[mid]; !ok {
			movies[mid] = &Item{
				Movie: popularMovie,
				Score: popularMovie.Popularity,
			}
		}

		movieGenre := popularMovie.Genres
		movieGenres := strings.Split(movieGenre, ",")

		for _, g := range movieGenres {
			if _, ok := preferenceSet[g]; ok {
				movies[mid].Score += 5
			}
		}
	}

	for _, rateMovie := range highRateMovies {
		mid := rateMovie.MovieID

		if _, ok := movies[mid]; !ok {
			movies[mid] = &Item{
				Movie: rateMovie,
				Score: rateMovie.Rate,
			}
		} else {
			movies[mid].Score += rateMovie.Rate
		}

		movieGenre := rateMovie.Genres
		movieGenres := strings.Split(movieGenre, ",")

		for _, g := range movieGenres {
			if _, ok := preferenceSet[g]; ok {
				movies[mid].Score += 5
			}
		}
	}

	rawItems := make([]*Item, 0, len(movies))
	for _, m := range movies {
		rawItems = append(rawItems, m)
	}
	//  sort
	sort.Slice(rawItems, func(i, j int) bool {
		return rawItems[i].Score < rawItems[j].Score
	})

	results := make([]infra.MovieModel, 0, len(rawItems))

	for _, i := range rawItems {
		results = append(results, i.Movie)
	}

	return results, nil
}

func containsGenre(genres []string, target string) bool {
	for _, genre := range genres {
		if strings.TrimSpace(genre) == target {
			return true
		}
	}
	return false
}
