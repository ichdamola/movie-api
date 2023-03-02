package app

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/ichdamola/movie-api/models"
)

type handlers struct {
	redisClient *redis.Client
	db          *sql.DB
}

func (h handlers) getMoviesHandler(w http.ResponseWriter, r *http.Request) {

	movies, err := getMovies(h.redisClient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sort.Slice(movies, func(i, j int) bool {
		return movies[i].ReleaseDate.Before(movies[j].ReleaseDate)
	})

	var movieList []models.Movie
	for _, movie := range movies {
		commentCount, err := getCommentCount(h.redisClient, movie.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		movieList = append(movieList, models.Movie{
			Title:        movie.Title,
			OpeningCrawl: movie.OpeningCrawl,
			CommentCount: commentCount,
		})
	}

	json.NewEncoder(w).Encode(movieList)
}

func (h handlers) handleCommentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieID, err := strconv.Atoi(vars["movie"])

	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	if r.Method == "POST" {
		var comment models.Comment
		err := json.NewDecoder(r.Body).Decode(&comment)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		comment.MovieID = movieID
		comment.IPAddress = r.RemoteAddr
		comment.CreatedAt = time.Now().UTC()

		err = addComment(h.db, comment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	} else if r.Method == "GET" {
		comments, err := getComments(h.db, movieID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(comments)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h handlers) getMovieCharacterHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieID, err := strconv.Atoi(vars["movie_id"])
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	characterList, err := getCharacters(h.db, movieID, r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(characterList)
}
