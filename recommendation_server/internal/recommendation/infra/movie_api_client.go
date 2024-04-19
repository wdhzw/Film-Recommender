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
    ID                 uint      `json:"id"`
    MovieID            int64     `json:"movie_id"`
    Title              string    `json:"title"`
    Overview           string    `json:"overview"`
    Rate               float64   `json:"rate"`
    Popularity         float64   `json:"popularity"`
    Homepage           string    `json:"homepage"`
    PosterURI          string    `json:"poster_uri"`
    Actors             []byte    `json:"actors"`
    Director           string    `json:"director"`
    Writers            string    `json:"writers"`
    Genres             string    `json:"genres"`
    ProductionCountry  string    `json:"production_country"`
    Language           string    `json:"language"`
    ReleaseDate        time.Time `json:"release_date"`
    Duration           int       `json:"duration"`
    Keyword            string    `json:"keyword"`
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