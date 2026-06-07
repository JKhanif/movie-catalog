package repo

import (
	"context"
	"movie_catalog/model"
)

func (r *Repository) CreateReview(ctx context.Context, review model.Review) error {

	return nil
}

func (r *Repository) GetReviewsByMovieID(ctx context.Context, id int) ([]model.Review, error) {

	return nil, nil
}

// func (r *Repository) GetReviewByID(ctx context.Context, id int) (model.Review, error) {

// 	return model.Review{}, nil
// }

// func (r *Repository) UpdateReview(ctx context.Context, id int, name string) {

// }

func (r *Repository) DeleteReview(ctx context.Context, id int, name string) {

}
