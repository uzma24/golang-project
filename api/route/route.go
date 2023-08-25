package route

import (
	"database/sql"

	"uzma/go-backend-clean-architecture/api/controller"
	"uzma/go-backend-clean-architecture/repository"
	"uzma/go-backend-clean-architecture/usecase"

	redis "github.com/go-redis/redis/v8"
	echo "github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo, db *sql.DB, redisdb *redis.Options) {
	mr := repository.NewMovieRepository(db)
	mc := &controller.MovieController{
		MovieUsecase: usecase.NewMovieUsecase(mr, redisdb),
	}
	v1 := e.Group("/v1")
	{
		movies := v1.Group("/movie")
		{
			movies.GET("/:id", mc.Fetch)
		}
	}
}
