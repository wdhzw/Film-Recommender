// config/config.go
package config

import (
    //"log"
    "os"
)

type Config struct {
    MovieAPIURL string
    UserAPIURL  string
}

func GetConfig() Config {
    return Config{
        MovieAPIURL: getEnv("MOVIE_API_URL", "http://localhost:5001"),
        UserAPIURL:  getEnv("USER_API_URL", "http://localhost:5002"), // Default ports, adjust as necessary
    }
}

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}
