package model

import (
	"CS5224_ESRS/movie"
	"time"
)

type MovieModel struct {
	ID                uint      `gorm:"column:id;primaryKey"`
	MovieID           int64     `gorm:"column:movie_id" json:"movie_id"`
	Title             string    `gorm:"column:title" json:"title"`
	Overview          string    `gorm:"column:overview" json:"overview"`
	Rate              float32   `gorm:"column:rate" json:"rate"`
	Popularity        float32   `gorm:"column:popularity" json:"popularity"`
	Homepage          string    `gorm:"column:homepage" json:"homepage"`
	PosterURI         string    `gorm:"column:poster_uri" json:"poster_uri"`
	Actors            []byte    `gorm:"column:actors" json:"actors"`
	Director          string    `gorm:"column:director" json:"director"`
	Writers           string    `gorm:"column:writers" json:"writers"`
	Genres            string    `gorm:"column:genres" json:"genres"`
	ProductionCountry string    `gorm:"column:production_country" json:"production_country"`
	Language          string    `gorm:"column:language" json:"language"`
	ReleaseDate       time.Time `gorm:"column:release_date" json:"release_date"`
	Duration          int       `gorm:"column:duration" json:"duration"`
	Keyword           string    `gorm:"column:keyword" json:"keyword"`
}

func (mm *MovieModel) TableName() string {
	return "movies"
}

type MovieProxy struct{}

func InitMovie() *MovieProxy {
	return &MovieProxy{}
}

func (p *MovieProxy) GetMovieModelById(movieId int64) (*MovieModel, error) {
	var movie *MovieModel
	movie_server.Mysql.Where("movie_id = ?", movieId).First(&movie)
	return movie, nil
}
