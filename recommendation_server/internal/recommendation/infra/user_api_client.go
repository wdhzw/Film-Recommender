package infra

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

type UserApiClient struct {
    BaseURL    string
    HTTPClient *http.Client
}

type UserTable struct {
    UserID          string   `json:"user_id"`
    UserName        string   `json:"user_name"`
    Email           string   `json:"email"`
    CreateTime      int64    `json:"create_time"`
    UpdateTime      int64    `json:"update_time"`
    PreferredGenre  []string `json:"preferred_genre,omitempty"`
}

func NewUserApiClient(baseURL string) *UserApiClient {
    return &UserApiClient{
        BaseURL: baseURL,
        HTTPClient: &http.Client{
            Timeout: time.Second * 30,
        },
    }
}

func (c *UserApiClient) FetchUserByEmail(email string) (*UserTable, error) {
    // Prepare the request body
    body := map[string]string{"email": email}
    jsonBody, _ := json.Marshal(body)

    // Send the request
    resp, err := c.HTTPClient.Post(c.BaseURL+"/get_user_by_email", "application/json", bytes.NewBuffer(jsonBody))
    if err != nil {
        return nil, fmt.Errorf("error fetching user by email: %v", err)
    }
    defer resp.Body.Close()

    // Parse the response
    var user UserTable
    if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
        return nil, fmt.Errorf("error decoding user data: %v", err)
    }

    return &user, nil
}