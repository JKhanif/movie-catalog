package repo

import (
	"context"
	"errors"
	"fmt"
	"movie_catalog/model"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetAllMovies(ctx context.Context) {

}

func (r *Repository) GetTopMovies(ctx context.Context, minReviews int) ([]model.MovieWithStats, error) {
	var movies []model.MovieWithStats

	rows, err := r.db.Query(ctx, queryGetTopMovies, minReviews)
	if err != nil {
		return nil, fmt.Errorf("failed to get top movies: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var m model.MovieWithStats
		err := rows.Scan(&m.ID, &m.Title, &m.Year, &m.Description,
			&m.Director, &m.CreatedAt, &m.UpdatedAt, &m.AvgRating, &m.ReviewCount)
		if err != nil {
			return nil, fmt.Errorf("failed to scan movie: %w", err)
		}

		movies = append(movies, m)
	}

	return movies, nil
}

func (r *Repository) GetMovieByID(ctx context.Context, id int) (model.MovieWithStats, error) {
	var m model.MovieWithStats

	err := r.db.QueryRow(ctx, queryGetMovieByID, id).Scan(
		&m.ID, &m.Title, &m.Year, &m.Description, &m.Director,
		&m.CreatedAt, &m.UpdatedAt, &m.AvgRating, &m.ReviewCount)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return m, ErrMovieNotFound
		}
		return m, fmt.Errorf("failed to get movie by ID: %w", err)
	}

	genres, err := r.GetGenresByMovieID(ctx, id)
	if err != nil {
		return m, fmt.Errorf("failed to get genres for movie: %w", err)
	}

	actors, err := r.GetActorsByMovieID(ctx, id)
	if err != nil {
		return m, fmt.Errorf("failed to get actors: %w", err)
	}

	m.Actors = actors
	m.Genres = genres

	return m, nil
}

// Admin handlers

func (r *Repository) CreateMovie(ctx context.Context, m model.Movie) (int, error) {
	var id int

	err := r.db.QueryRow(ctx, queryCreateMovie, m.Title, m.Year, m.Description, m.Director).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create movie: %w", err)
	}

	return id, nil
}

func (r *Repository) UpdateMovie(ctx context.Context, id int, m model.Movie) error {
	res, err := r.db.Exec(ctx, queryUpdateMovie, m.Title, m.Year, m.Description, m.Director, id)
	if err != nil {
		return fmt.Errorf("failed to update movie: %w", err)
	}

	rows := res.RowsAffected()
	if rows == 0 {
		return ErrMovieNotFound
	}

	return nil
}

func (r *Repository) DeleteMovie(ctx context.Context, id int) error {
	res, err := r.db.Exec(ctx, queryDeleteMovie, id)
	if err != nil {
		return fmt.Errorf("failed to delete movie: %w", err)
	}

	rows := res.RowsAffected()
	if rows == 0 {
		return ErrMovieNotFound
	}

	return nil
}
