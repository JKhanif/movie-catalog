package repo

import (
	"context"
	"errors"
	"fmt"
	"movie_catalog/model"
	"strings"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetAllMovies(ctx context.Context, genre string, year int, minRating float64, search string, sort string, order string, limit int, offset int) ([]model.MovieWithStats, error) {
	validSorts := map[string]bool{"year": true, "rating": true, "created_at": true}
	validOrders := map[string]bool{"asc": true, "desc": true}

	if !validSorts[sort] {
		sort = "created_at"
	}
	if !validOrders[order] {
		order = "desc"
	}

	var sortColumn string
	switch sort {
	case "rating":
		sortColumn = "avg_rating"
	case "year":
		sortColumn = "m.year"
	default:
		sortColumn = "m.created_at"
	}

	direction := "DESC"
	if order == "asc" {
		direction = "ASC"
	}

	query := queryGetAllMovies + fmt.Sprintf(" ORDER BY %s %s LIMIT $5 OFFSET $6", sortColumn, direction)

	rows, err := r.db.Query(ctx, query, genre, year, minRating, search, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get movies: %w", err)
	}
	defer rows.Close()

	var movies []model.MovieWithStats
	for rows.Next() {
		var m model.MovieWithStats
		err := rows.Scan(&m.ID, &m.Title, &m.Year, &m.Description,
			&m.Director, &m.CreatedAt, &m.UpdatedAt, &m.AvgRating, &m.ReviewCount)
		if err != nil {
			return nil, fmt.Errorf("failed to scan movie: %w", err)
		}
		m.Genres = []model.Genre{}
		m.Actors = []model.Actor{}
		movies = append(movies, m)
	}

	return movies, nil
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

		m.Genres = []model.Genre{}
		m.Actors = []model.Actor{}
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

func (r *Repository) UpdateMovie(ctx context.Context, id int, req model.UpdateMovieRequest) error {
	var sets []string
	var args []any
	argIdx := 1

	if req.Title != nil {
		sets = append(sets, fmt.Sprintf("title = $%d", argIdx))
		args = append(args, *req.Title)
		argIdx++
	}
	if req.Year != nil {
		sets = append(sets, fmt.Sprintf("year = $%d", argIdx))
		args = append(args, *req.Year)
		argIdx++
	}
	if req.Description != nil {
		sets = append(sets, fmt.Sprintf("description = $%d", argIdx))
		args = append(args, *req.Description)
		argIdx++
	}
	if req.Director != nil {
		sets = append(sets, fmt.Sprintf("director = $%d", argIdx))
		args = append(args, *req.Director)
		argIdx++
	}

	if len(sets) == 0 {
		return ErrNoFieldsToUpdate
	}

	sets = append(sets, "updated_at = NOW()")
	args = append(args, id)

	query := fmt.Sprintf("UPDATE movies SET %s WHERE id = $%d", strings.Join(sets, ", "), argIdx)

	res, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update movie: %w", err)
	}

	if res.RowsAffected() == 0 {
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
