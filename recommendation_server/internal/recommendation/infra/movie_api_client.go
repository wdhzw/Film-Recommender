package infra

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// MovieApiClient represents a client for making API calls to fetch movie data.
type MovieApiClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

// Pagination defines the structure for API response pagination details.
type Pagination struct {
	CurrentPage int `json:"current_page"`
	TotalNumber int `json:"total_number"`
	TotalPages  int `json:"total_pages"`
}

// PopularMovieModel defines the structure for movies fetched from the popular movies endpoint.
type PopularMovieModel struct {
	MovieID     int64     `json:"movie_id"`
	Title       string    `json:"title"`
	Overview    string    `json:"overview"`
	Popularity  float64   `json:"popularity"`
	Genres      []string  `json:"genres"`
	ReleaseDate time.Time `json:"release_date"`
	Duration    int       `json:"duration"`
	Keywords    []string  `json:"keywords"`
	PosterURL   string    `json:"poster_uri"`
}

// HighRateMovieModel defines the structure for movies fetched from the high-rate movies endpoint.
type HighRateMovieModel struct {
	MovieID     int64     `json:"movie_id"`
	Title       string    `json:"title"`
	Rate        float64   `json:"rate"`
	Genres      []string  `json:"genres"`
	ReleaseDate time.Time `json:"release_date"`
	Duration    int       `json:"duration"`
	Keywords    []string  `json:"keywords"`
	PosterURL   string    `json:"poster_uri"`
}

// PopularMovieResponse wraps the response from the popular movies endpoint.
type PopularMovieResponse struct {
	Content    []PopularMovieModel `json:"content"`
	Pagination Pagination          `json:"pagination"`
}

// HighRateMovieResponse wraps the response from the high-rate movies endpoint.
type HighRateMovieResponse struct {
	Content    []HighRateMovieModel `json:"content"`
	Pagination Pagination           `json:"pagination"`
}

// NewMovieApiClient creates a new instance of MovieApiClient.
func NewMovieApiClient(baseURL string) *MovieApiClient {
	return &MovieApiClient{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

func (c *MovieApiClient) GetPopularMovies(pageNumber int) (PopularMovieResponse, error) {
	url := fmt.Sprintf("%s/popular?page=%d", c.BaseURL, pageNumber)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return PopularMovieResponse{}, fmt.Errorf("error fetching popular movies: %v", err)
	}
	defer resp.Body.Close()

	var response PopularMovieResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return PopularMovieResponse{}, fmt.Errorf("error decoding popular movies: %v", err)
	}

	return response, nil
}

func (c *MovieApiClient) GetHighRateMovies(pageNumber int) (HighRateMovieResponse, error) {
	url := fmt.Sprintf("%s/top_rate?page=%d", c.BaseURL, pageNumber)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return HighRateMovieResponse{}, fmt.Errorf("error fetching high-rate movies: %v", err)
	}
	defer resp.Body.Close()

	var response HighRateMovieResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return HighRateMovieResponse{}, fmt.Errorf("error decoding high-rate movies: %v", err)
	}

	return response, nil
}
