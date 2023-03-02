package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"github.com/ichdamola/movie-api/models"
)

func getMovies(redisClient *redis.Client) ([]models.MovieDetails, error) {

	movieData, err := redisClient.Get("movies").Result()
	if err == redis.Nil {
		// Cache miss, fetch data from swapi.dev
		resp, err := http.Get("https://swapi.dev/api/films/")
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var data struct {
			Results []models.MovieDetails `json:"results"`
		}

		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			return nil, err
		}

		movieData, err := json.Marshal(data.Results)
		if err != nil {
			return nil, err
		}

		redisClient.Set("movies", movieData, 1*time.Hour)

	} else if err != nil {
		return nil, err
	}

	var movies []models.MovieDetails
	err = json.Unmarshal([]byte(movieData), &movies)
	if err != nil {
		return nil, err
	}

	return movies, nil

}

func getCommentCount(redisClient *redis.Client, movieID int) (int, error) {
	commentCount, err := redisClient.Get(fmt.Sprintf("comment_count:%d", movieID)).Int()
	if err == redis.Nil {
		// Cache miss, fetch data from
		return commentCount, nil
	}

	return 1, nil
}
