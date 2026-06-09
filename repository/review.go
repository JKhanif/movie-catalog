package repo

import (
	"context"
	"fmt"
	"movie_catalog/model"
)

func (r *Repository) CreateReview(ctx context.Context, movieID int, author string, rating int, text string) (int, error) {
	var id int

	err := r.db.QueryRow(ctx, queryCreateReview, movieID, author, rating, text).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create review: %w", err)
	}

	return id, nil
}

func (r *Repository) GetReviewsByMovieID(ctx context.Context, id int, limit int, offset int) ([]model.Review, error) {
	var review []model.Review

	rows, err := r.db.Query(ctx, queryGetReviewsByMovieID, id, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query reviews: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var rew model.Review
		err := rows.Scan(&rew.ID, &rew.MovieID, &rew.Author, &rew.Rating, &rew.Text, &rew.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan review: %w", err)
		}

		review = append(review, rew)
	}

	return review, nil
}

// func (r *Repository) GetReviewByID(ctx context.Context, id int) (model.Review, error) {

// 	return model.Review{}, nil
// }

// func (r *Repository) UpdateReview(ctx context.Context, id int, name string) {

// }

func (r *Repository) DeleteReview(ctx context.Context, id int) error {
	res, err := r.db.Exec(ctx, queryDeleteReview, id)
	if err != nil {
		return fmt.Errorf("failed to delete review: %w", err)
	}

	rows := res.RowsAffected()
	if rows == 0 {
		return ErrReviewNotFound
	}

	return nil
}
