package controller

import (
	"net/http"
	"strconv"
	"uzma/go-backend-clean-architecture/domain"

	"github.com/labstack/echo/v4"
)

type MovieController struct {
	MovieUsecase domain.MovieUsecase
}

func (mc *MovieController) Fetch(c echo.Context) error {
	movieId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return err
	}

	movie, err := mc.MovieUsecase.FetchByMovieID(c, int64(movieId))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return err
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Data fetched successfully!!",
		Data:    movie,
	})

	return nil
}
