package model

import "time"

type Work struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Maker       Maker     `json:"maker"`
	Author      string    `json:"author"`
	Price       int       `json:"price"`
	Discount    int       `json:"discount"`
	DL          int       `json:"dl"`
	URL         string    `json:"url"`
	RatingStar  int       `json:"star"`
	RatingTotal int       `json:"rating"`
	FetchedAt   time.Time `json:"fetchedAt"`
}
