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
        MovieAPIURL: getEnv("MOVIE_API_URL", "http://cs5224-movie-service-env.eba-ptufih3p.us-east-1.elasticbeanstalk.com/movie_server"),
        UserAPIURL:  getEnv("USER_API_URL", "http://ec2-44-217-97-83.compute-1.amazonaws.com:8080/api"),
    }
}

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}