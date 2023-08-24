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

// type MovieController struct {
// 	MovieUsecase usecase.MovieUsecaseService
// }

// func NewMovieController(movieSrv usecase.MovieUsecaseService) *MovieController {
// 	return &MovieController{
// 		MovieUsecase: movieSrv,
// 	}
// }

// func (m *MovieController) Fetch(c echo.Context) {

// 	movieId, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {

// 	}

// 	movie, err := m.MovieUsecase.Get(c.Request().Context(), movieId)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, err)
// 		return
// 	}

// 	c.JSON(http.StatusOK, domain.SuccessResponse{
// 		Message: "Data fetched successfully!!",
// 		Data:    movie,
// 	})
// }
