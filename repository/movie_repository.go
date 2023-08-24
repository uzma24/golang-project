package repository

import (
	"database/sql"
	"log"
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
	log.Println("movie id in repo: ", movieID)

	query := "SELECT id, title, poster_path, language, overview, releasedate FROM movie WHERE id = ?"

	var movie domain.Movie
	err := mr.database.QueryRow(query, movieID).Scan(&movie.ID, &movie.Title, &movie.PosterPath, &movie.Language, &movie.Overview, &movie.ReleaseDate)
	log.Printf("movie in repo: %+v ", movie)
	if err != nil {
		log.Println(err.Error())
	}

	// err := mr.database.QueryRow(query, movieID).Scan(&movie)
	// row, err:= mr.database.Query(query)
	//row, err := mr.database.Query("SELECT * FROM movie WHERE id = ?", movieID)
	//log.Println("movie row in repo: ", row)

	//err1 := json.Unmarshal(row, &movie)

	if err != nil {
		return domain.Movie{}, err
	}

	return movie, nil
}
