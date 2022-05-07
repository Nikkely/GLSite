package model

type Work struct {
	ID          string
	Name        string
	MakerName   string // TODO: create Maker model
	Author      string
	Price       int
	Discount    int
	DL          int
	URL         string
	RatingStar  int
	RatingTotal int
	Label       string
}
