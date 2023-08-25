package repository

import (
	"database/sql"
	"uzma/go-backend-clean-architecture/domain"

	echo "github.com/labstack/echo/v4"
)

type movieRepository struct {
	database *sql.DB
}

func NewMovieRepository(db *sql.DB) domain.MovieRepository {
	return &movieRepository{
		database: db,
	}
}

func (mr *movieRepository) FetchByMovieID(e echo.Context, movieID int64) (domain.Movie, error) {
	query := "SELECT id, title, poster_path, language, overview, releasedate FROM movie WHERE id = ?"

	var movie domain.Movie
	err := mr.database.QueryRow(query, movieID).Scan(&movie.ID, &movie.Title, &movie.PosterPath, &movie.Language, &movie.Overview, &movie.ReleaseDate)

	if err != nil {
		return domain.Movie{}, err
	}

	return movie, nil
}
