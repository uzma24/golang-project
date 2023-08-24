package domain

import (
	"time"

	echo "github.com/labstack/echo/v4"
)

type Movie struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	PosterPath  string    `json:"poster_path"`
	Language    string    `json:"language" default:""`
	Overview    string    `json:"overview" default:""`
	ReleaseDate time.Time `json:"releasedate"`
}

type MovieRepository interface {
	FetchByMovieID(e echo.Context, movieID int64) (Movie, error)
}

type MovieUsecase interface {
	FetchByMovieID(e echo.Context, movieID int64) (Movie, error)
}
