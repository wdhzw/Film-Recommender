package model

import (
	"CS5224_ESRS/movie/infra"
	"fmt"
	"time"
)

type MovieModel struct {
	ID                uint      `gorm:"column:id;primaryKey"`
	MovieID           int64     `gorm:"column:movie_id" json:"movie_id"`
	Title             string    `gorm:"column:title" json:"title"`
	Overview          string    `gorm:"column:overview" json:"overview"`
	Rate              float64   `gorm:"column:rate" json:"rate"`
	Popularity        float64   `gorm:"column:popularity" json:"popularity"`
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
	err := infra.Mysql.Model(&MovieModel{}).Where("movie_id = ?", movieId).Find(&movie).Error
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (p *MovieProxy) GetPopularMovies(page int) ([]*MovieModel, error) {
	movies := make([]*MovieModel, 0)
	offset := (page - 1) * 10

	err := infra.Mysql.Model(&MovieModel{}).Order("popularity DESC").
		Limit(10).
		Offset(offset).
		Find(&movies).Error

	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (p *MovieProxy) CountTotalMovies() (total int64, err error) {
	err = infra.Mysql.Model(&MovieModel{}).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (p *MovieProxy) GetHighRateMovies(page int) ([]*MovieModel, error) {
	movies := make([]*MovieModel, 0)
	offset := (page - 1) * 10

	err := infra.Mysql.Model(&MovieModel{}).Order("rate DESC").
		Limit(10).
		Offset(offset).
		Find(&movies).Error

	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (p *MovieProxy) UpdateMovies(movieId int64, popularity, rate float64) error {
	err := infra.Mysql.Model(&MovieModel{}).
		Where("movie_id = ? ", movieId).
		UpdateColumns(map[string]interface{}{
			"popularity": popularity,
			"rate":       rate,
		}).Error

	return err
}

func (p *MovieProxy) CreateMovie(model *MovieModel) error {
	return infra.Mysql.Model(&MovieModel{}).Create(model).Error
}

func (p *MovieProxy) SearchMovie(word string) ([]*MovieModel, error) {
	movies := make([]*MovieModel, 0)
	err := infra.Mysql.Model(&MovieModel{}).
		Where("title LIKE ?", fmt.Sprint("%s%", word)).
		Limit(50).
		Find(&movies).Error

	if err != nil {
		return nil, err
	}
	return movies, nil
}
