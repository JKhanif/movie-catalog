package repo

import (
	"context"
	"fmt"
	"movie_catalog/model"
)

func (r *Repository) GetAllGenres(ctx context.Context) ([]model.Genre, error) {
	var g []model.Genre

	rows, err := r.db.Query(ctx, queryGetAllGenres)
	if err != nil {
		return nil, fmt.Errorf("failed to get all genres: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var genre model.Genre
		err := rows.Scan(&genre.ID, &genre.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan genre: %w", err)
		}

		g = append(g, genre)
	}

	return g, nil
}

func (r *Repository) GetGenresByMovieID(ctx context.Context, id int) ([]model.Genre, error) {

	return nil, nil
}

// func (r *Repository) GetGenreByID(ctx context.Context, id int) (model.Genre, error) {
// 	var g model.Genre

// 	err := r.db.QueryRow(ctx, queryGetGenreByID, id).Scan(&g.ID, &g.Name)
// 	if err != nil {
// 		if errors.Is(err, pgx.ErrNoRows) {
// 			return g, ErrGenreNotFound
// 		}
// 		return g, fmt.Errorf("failed to get genre by ID: %w", err)
// 	}

// 	return g, nil
// }

// Admin handlers

func (r *Repository) CreateGenre(ctx context.Context, g model.Genre) error {
	_, err := r.db.Exec(ctx, queryCreateGenre, g.Name)
	if err != nil {
		return fmt.Errorf("failed to create genre: %w", err)
	}

	return nil
}

// func (r *Repository) UpdateGenre(ctx context.Context) {

// }

func (r *Repository) DeleteGenre(ctx context.Context, id int) error {
	res, err := r.db.Exec(ctx, queryDeleteGenre, id)
	if err != nil {
		return fmt.Errorf("failed to delete genre: %w", err)
	}

	rows := res.RowsAffected()
	if rows == 0 {
		return ErrGenreNotFound
	}

	return nil
}
