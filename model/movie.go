package model

import "time"

type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Year        int       `json:"year"`
	Description string    `json:"description"`
	Director    string    `json:"director"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateMovieRequest struct {
	Title       *string `json:"title,omitempty"`
	Year        *int    `json:"year,omitempty"`
	Description *string `json:"description,omitempty"`
	Director    *string `json:"director,omitempty"`
}
