package usecase

import (
	"uzma/go-backend-clean-architecture/domain"

	echo "github.com/labstack/echo/v4"
)

type movieUsecase struct {
	movieRepository domain.MovieRepository
}

func NewMovieUsecase(movieRepository domain.MovieRepository) domain.MovieUsecase {
	return &movieUsecase{
		movieRepository: movieRepository,
	}
}

func (mu *movieUsecase) FetchByMovieID(c echo.Context, movieID int64) (domain.Movie, error) {
	return mu.movieRepository.FetchByMovieID(c, movieID)
}

// type movieUsecase struct {
// 	movieRepository repository.MovieRepository
// 	// contextTimeout  time.Duration
// }

// type MovieUsecaseService interface {
// 	Get(c context.Context, id int) (domain.Movie, error)
// }

// func NewMovieUsecase(movieRepository repository.MovieRepository) *movieUsecase {
// 	return &movieUsecase{
// 		movieRepository: movieRepository,
// 	}
// }

// func (mu *movieUsecase) Get(c context.Context, id int) (domain.Movie, error) {
// 	ctx := context.Background()
// 	defer ctx.Done()
// 	movie, err := mu.movieRepository.Get(ctx, id)
// 	if err != nil {
// 		return domain.Movie{}, err
// 	}

// 	return *movie, nil
// }
