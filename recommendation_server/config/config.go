package config

import (
    "os"
)

type Config struct {
    MovieAPIURL string
    UserAPIURL  string
}

func GetConfig() Config {
    return Config{
        MovieAPIURL: getEnv("MOVIE_API_URL", "http://cs5224-movie-service.us-east-1.elasticbeanstalk.com/movie_server"),
        UserAPIURL:  getEnv("USER_API_URL", "http://ec2-44-217-97-83.compute-1.amazonaws.com:8080/api"),
        // MovieAPIURL: getEnv("MOVIE_API_URL", "http://localhost:5001"),
        // UserAPIURL:  getEnv("USER_API_URL", "http://localhost:8088/api"),
    }
}

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}