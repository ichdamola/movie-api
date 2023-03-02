package main

import (
	"database/sql"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/ichdamola/movie-api/app"
)

func main() {

	r := mux.NewRouter()

	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=comments sslmode=disable")
	if err != nil {
		panic(err)
	}

	app.Run(r, redisClient, db)
}
