package infra

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
    "recommendation_server/pkg/entity"
)

type MovieApiClient struct {
    BaseURL    string
    HTTPClient *http.Client
}

func NewMovieApiClient(baseURL string) *MovieApiClient {
    return &MovieApiClient{
        BaseURL: baseURL,
        HTTPClient: &http.Client{
            Timeout: time.Second * 30, // Adjust the timeout as necessary
        },
    }
}

// FetchMovies fetches movies based on the provided preferences.
// This is a simplified example and might need adjustments based on your actual API.
func (c *MovieApiClient) FetchMovies(preferences UserPreferences) ([]entity.Movie, error) {
    // Construct the URL with query parameters based on preferences.
    // This is an example; your implementation details might vary.
    url := fmt.Sprintf("%s/movies?genre=%s", c.BaseURL, preferences.FavoriteGenre)

    resp, err := c.HTTPClient.Get(url)
    if err != nil {
        return nil, fmt.Errorf("error fetching movies: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
    }

    var movies []entity.Movie
    if err := json.NewDecoder(resp.Body).Decode(&movies); err != nil {
        return nil, fmt.Errorf("error decoding movies: %v", err)
    }

    return movies, nil
}
