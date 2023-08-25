package usecase

import (
	"encoding/json"
	"log"
	"strconv"
	"time"
	"uzma/go-backend-clean-architecture/domain"

	redis "github.com/go-redis/redis/v8"
	echo "github.com/labstack/echo/v4"
)

type movieUsecase struct {
	movieRepository domain.MovieRepository
	redisOptions    *redis.Options
}

func NewMovieUsecase(movieRepository domain.MovieRepository, redisOp *redis.Options) domain.MovieUsecase {
	return &movieUsecase{
		movieRepository: movieRepository,
		redisOptions:    redisOp,
	}
}

func (mu *movieUsecase) FetchByMovieID(c echo.Context, movieID int64) (domain.Movie, error) {
	// get from redis
	//if resp is nil then call repo
	ctx := c.Request().Context()
	redisClient := redis.NewClient(mu.redisOptions)
	movieResultJSON, err := redisClient.Get(ctx, strconv.FormatInt(movieID, 10)).Result()
	if err != nil {
		log.Println("Error getting Movie value from Redis:: ", err)
	}

	var mov *domain.Movie
	err = json.Unmarshal([]byte(movieResultJSON), &mov)
	if err != nil {
		log.Println("Err Unmarshalling Movie Result from Redis:: ", err)
	}

	var movie domain.Movie

	movie, err = mu.movieRepository.FetchByMovieID(c, movieID)
	if err != nil {
		log.Println("Error fetching from SQL DB :: ", err)
	}

	// now set the data into redis with some ttl
	movieJSON, err := json.Marshal(movie)
	if err != nil {
		log.Println("Error Marshalling movie!!")
	}
	err = redisClient.Set(ctx, strconv.FormatInt(movieID, 10), movieJSON, time.Duration(time.Duration.Hours(24))).Err()
	if err == redis.Nil {
		log.Println("Not able to set redis key!!", err)
	}

	return movie, nil
}
