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
type UserApiResponse struct {
    Data       UserTable  `json:"data"`
    Error      *string    `json:"error"`       // Using a pointer to handle null
    StatusCode int        `json:"status_code"`
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
    var response UserApiResponse
    if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
        return nil, fmt.Errorf("error decoding user data: %v", err)
    }

    // Handle potential API-reported errors
    if response.Error != nil {
        return nil, fmt.Errorf("API error: %s", *response.Error)
    }

    // Check the status code for non-success cases
    if response.StatusCode != 0 { // Assuming 0 means success
        return nil, fmt.Errorf("API returned non-success status code: %d", response.StatusCode)
    }
    //fmt.Printf("Fetched User: %+v\n", response.Data)
    return &response.Data, nil
}
