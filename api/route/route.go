package route

import (
	"database/sql"

	"uzma/go-backend-clean-architecture/api/controller"
	"uzma/go-backend-clean-architecture/repository"
	"uzma/go-backend-clean-architecture/usecase"

	echo "github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo, db *sql.DB) {
	mr := repository.NewMovieRepository(db)
	mc := &controller.MovieController{
		MovieUsecase: usecase.NewMovieUsecase(mr),
	}
	v1 := e.Group("/v1")
	{
		movies := v1.Group("/movie")
		{
			movies.GET("/:id", mc.Fetch)
		}
	}
}
