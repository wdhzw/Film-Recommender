package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"io"
	"strconv"
	"time"

	//"github.com/aws/aws-lambda-go/lambda"
	"github.com/ryanbradynd05/go-tmdb"
	"net/http"
)

var (
	tmdbAPI = tmdb.Init(tmdb.Config{
		APIKey:   "4f40401acacceee6b426290d18f6f38a",
		Proxies:  nil,
		UseProxy: false,
	})
)

type MovieContent struct {
	Content struct {
		MovieId int64 `json:"movie_id"`
	} `json:"content"`
}

type UpdateRequest struct {
	MovieId    string `form:"movie_id"`
	Rate       string `form:"rate"`
	Popularity string `form:"popularity"`
}

type Actor struct {
	Name       string `json:"name" form:"movie_id"`
	Character  string `json:"character" form:"movie_id"`
	ProfileUri string `json:"profile_uri" form:"movie_id"`
}

type CreateMovieRequest struct {
	MovieID             int64     `form:"movie_id"`
	Title               string    `form:"title"`
	Overview            string    `form:"overview"`
	Rate                float64   `form:"rate"`
	Popularity          float64   `form:"popularity"`
	Homepage            string    `form:"homepage"`
	PosterURI           string    `form:"poster_uri"`
	Actors              []*Actor  `form:"actors"`
	Director            string    `form:"director"`
	Writers             []string  `form:"writers"`
	Genres              []string  `form:"genres"`
	ProductionCountries []string  `form:"production_countries"`
	Languages           []string  `form:"languages"`
	ReleaseDate         time.Time `form:"release_date"`
	Duration            int       `form:"duration"`
	Keywords            []string  `form:"keywords"`
}

func fetchMovies() {

	for i := 0; i < 10; i++ {
		start, end := i*50+1, (i+1)*50

		go func(start, end int) {

			for j := start; j <= end; j++ {
				resp, err := tmdbAPI.GetMoviePopular(map[string]string{
					"page": string(rune(j)),
				})

				if err != nil {
					fmt.Errorf("[GetMoviePopular] failed, err:%s", err)
					continue
				}

				movies := resp.Results

				for _, movie := range movies {
					movieId := movie.ID
					movieDetail, err := tmdbAPI.GetMovieInfo(movieId, nil)
					if err != nil {
						fmt.Errorf("[GetMovieInfo] failed, err:%s", err)
						continue
					}

					resp, err := http.Get(fmt.Sprintf("http://localhost:8080/movie_server/%d", movieId))
					if err != nil {
						fmt.Errorf("[GetMovieInfo] failed, err:%s", err)
						continue
					}
					defer resp.Body.Close()

					// 读取响应体
					body, err := io.ReadAll(resp.Body)
					if err != nil {
						fmt.Errorf("[movie server] Get movie api failed, err:%s", err)
						continue
					}

					content := &MovieContent{}
					_ = json.Unmarshal(body, content)

					if content.Content.MovieId == 0 {
						//id, _ := strconv.ParseInt(movieId, 10, 64)
						createReq := &CreateMovieRequest{
							MovieID:    int64(movieId),
							Title:      movie.Title,
							Overview:   movie.Overview,
							Rate:       float64(movie.VoteAverage),
							Popularity: float64(movie.Popularity),
							PosterURI:  movie.PosterPath,
						}
						homepage := movieDetail.Homepage
						if homepage == "" {
							homepage = "https://themoviedb.org/movie/" + string(movieId)
						}
						createReq.Homepage = homepage

						// actors
						creditResp, err := tmdbAPI.GetMovieCredits(movieId, nil)
						if err != nil {
							fmt.Errorf("[movie server] Get Movie Credit api failed, err:%s", err)
							continue
						}

						actors := make([]*Actor, 0)
						for i, cast := range creditResp.Cast {
							if i > 9 {
								break
							}
							actor := &Actor{
								Name:       cast.Name,
								Character:  cast.Character,
								ProfileUri: cast.ProfilePath,
							}
							actors = append(actors, actor)
						}

						createReq.Actors = actors

						// director && writer
						director := ""
						writers := make([]string, 0)
						for _, crew := range creditResp.Crew {
							if crew.Job == "Director" {
								director = crew.Name
							} else if crew.Job == "Story" || crew.Job == "Writer" || crew.Job == "Novel" || crew.Job == "Screenplay" {
								writers = append(writers, crew.Name)
							}

						}
						createReq.Director = director
						createReq.Writers = writers

						// genre
						genres := make([]string, 0)
						for _, genre := range movieDetail.Genres {
							genres = append(genres, genre.Name)
						}
						createReq.Genres = genres

						// production countries
						productCountries := make([]string, 0)
						for _, pc := range movieDetail.ProductionCountries {
							productCountries = append(productCountries, pc.Name)
						}
						createReq.ProductionCountries = productCountries

						// language
						languages := make([]string, 0)
						for _, l := range movieDetail.SpokenLanguages {
							languages = append(languages, l.Name)
						}
						createReq.Languages = languages
						// release_date
						releaseDate := movieDetail.ReleaseDate
						t, _ := time.Parse("2006-01-02", releaseDate)
						createReq.ReleaseDate = t
						// time duration, unit is minutes
						duration := movieDetail.Runtime
						createReq.Duration = int(duration)
						// keywords
						keywordResp, err := tmdbAPI.GetMovieKeywords(movieId, nil)
						if err != nil {
							fmt.Errorf("[movie server] Get Movie Keywords api failed, err:%s", err)
							continue
						}
						keywords := make([]string, 0)
						for _, w := range keywordResp.Keywords {
							keywords = append(keywords, w.Name)
						}
						createReq.Keywords = keywords

						createReqStr, _ := json.Marshal(createReq)
						_, err = http.Post("http://localhost:8080/movie_server/create", "application/json", bytes.NewBuffer(createReqStr))
						if err != nil {
							fmt.Errorf("[movie server] Create Movie Api Failed, err:%s", err)
							continue
						}
					} else {
						updateReq := &UpdateRequest{
							MovieId:    strconv.Itoa(movieId),
							Popularity: strconv.FormatFloat(float64(movie.Popularity), 'f', -1, 64),
							Rate:       strconv.FormatFloat(float64(movie.VoteAverage), 'f', -1, 64),
						}
						updateReqStr, _ := json.Marshal(updateReq)
						_, err := http.Post("http://localhost:8080/movie_server/update", "application/json", bytes.NewBuffer(updateReqStr))
						if err != nil {
							fmt.Errorf("[movie server] Update Movie Api Failed, err:%s", err)
							continue
						}
					}

				}

			}
		}(start, end)
	}

}

func main() {
	lambda.Start(fetchMovies)
	//fetchMovies()
}
