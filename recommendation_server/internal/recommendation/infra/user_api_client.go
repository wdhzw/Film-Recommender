package infra

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

type UserApiClient struct {
    BaseURL    string
    HTTPClient *http.Client
}

type UserPreferences struct {
    // Fields that match the JSON structure of the response from the user service
		FavoriteGenre string `json:"favorite_genre"`
}

func NewUserApiClient(baseURL string) *UserApiClient {
    return &UserApiClient{
        BaseURL: baseURL,
        HTTPClient: &http.Client{
            Timeout: time.Second * 30, // or another appropriate timeout
        },
    }
}

func (c *UserApiClient) FetchUserPreferences(userID string) (*UserPreferences, error) {
    // Construct the request URL
    url := fmt.Sprintf("%s/user/%s/preferences", c.BaseURL, userID)

    // Execute the request
    resp, err := c.HTTPClient.Get(url)
    if err != nil {
        return nil, fmt.Errorf("error fetching user preferences: %v", err)
    }
    defer resp.Body.Close()

    // Check the response status code
    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
    }

    // Decode the response body into the UserPreferences struct
    var preferences UserPreferences
    if err := json.NewDecoder(resp.Body).Decode(&preferences); err != nil {
        return nil, fmt.Errorf("error decoding user preferences: %v", err)
    }

    return &preferences, nil
}
