package models

import "time"

type (
	MovieDetails struct {
		ID           int       `json:"id"`
		Title        string    `json:"title"`
		OpeningCrawl string    `json:"opening_crawl"`
		ReleaseDate  time.Time `json:"release_date"`
	}
)

type Movie struct {
	Title        string `json:"title"`
	OpeningCrawl string `json:"opening_crawl"`
	CommentCount int    `json:"comment_count"`
}
