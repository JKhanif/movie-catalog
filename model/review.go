package model

import "time"

type Review struct {
	ID        int       `json:"id"`
	MovieID   int       `json:"movie_id"`
	Author    string    `json:"author"`
	Rating    int       `json:"rating"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
