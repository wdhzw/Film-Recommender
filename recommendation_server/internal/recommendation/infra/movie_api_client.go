package infra

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type MovieApiClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

type MovieModel struct {
	MovieId     int64     `json:"movie_id"`
	Title       string    `json:"title"`
	Popularity  float64   `json:"popularity"`
	Genres      []string  `json:"genres"`
	ReleaseDate time.Time `json:"release_date"`
	Duration    int       `json:"duration"`
	Keywords    []string  `json:"keywords"`
}

func NewMovieApiClient(baseURL string) *MovieApiClient {
	return &MovieApiClient{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

func (c *MovieApiClient) GetPopularMovies() ([]MovieModel, error) {
	// Send the request
	resp, err := c.HTTPClient.Get(c.BaseURL + "/popular")
	if err != nil {
		return nil, fmt.Errorf("error fetching popular movies: %v", err)
	}
	defer resp.Body.Close()

	// Parse the response
	var movies []MovieModel
	if err := json.NewDecoder(resp.Body).Decode(&movies); err != nil {
		return nil, fmt.Errorf("error decoding popular movies: %v", err)
	}

	return movies, nil
}

func (c *MovieApiClient) GetHighRateMovies() ([]MovieModel, error) {
	// Send the request
	resp, err := c.HTTPClient.Get(c.BaseURL + "/top_rate")
	if err != nil {
		return nil, fmt.Errorf("error fetching high-rate movies: %v", err)
	}
	defer resp.Body.Close()

	// Parse the response
	var movies []MovieModel
	if err := json.NewDecoder(resp.Body).Decode(&movies); err != nil {
		return nil, fmt.Errorf("error decoding high-rate movies: %v", err)
	}

	return movies, nil
}
