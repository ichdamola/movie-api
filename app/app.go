package app

import (
	"database/sql"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

func Run(r *mux.Router, redisClient *redis.Client, db *sql.DB) {

	handler := handlers{
		redisClient: redisClient,
		db:          db,
	}

	r.HandleFunc("/movies", handler.getMoviesHandler).Methods("GET")

	r.HandleFunc("/comments/{movie_id}", handler.handleCommentHandler).Methods("GET", "POST")

	r.HandleFunc("/characters/{movie_id}", handler.getMovieCharacterHandler).Methods("GET")

	http.ListenAndServe(":8000", r)

}
