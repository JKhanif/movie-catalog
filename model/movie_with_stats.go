package model

type MovieWithStats struct {
	Movie
	Genres      []Genre `json:"genres"`
	Actors      []Actor `json:"actors"`
	AvgRating   float64 `json:"avg_rating"`
	ReviewCount int     `json:"review_count"`
}
