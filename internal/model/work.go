package model

type Work struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	MakerName   string `json:"makerName"` // TODO: create Maker model
	Author      string `json:"author"`
	Price       int    `json:"price"`
	Discount    int    `json:"discount"`
	DL          int    `json:"dl"`
	URL         string `json:"url"`
	RatingStar  int    `json:"star"`
	RatingTotal int    `json:"rating"`
}
