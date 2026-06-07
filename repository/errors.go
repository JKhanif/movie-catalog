package repo

import "errors"

var ErrMovieNotFound = errors.New("movie not found")
var ErrGenreNotFound = errors.New("genre not found")
var ErrReviewNotFound = errors.New("review not found")
