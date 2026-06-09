package repo

import (
	"context"
	"fmt"
	"movie_catalog/model"
)

// import (
// 	"context"
// )

// func (r *Repository) GetAllActors(ctx context.Context) {

// }

func (r *Repository) GetActorsByMovieID(ctx context.Context, id int) ([]model.Actor, error) {
	var actors []model.Actor

	rows, err := r.db.Query(ctx, queryGetActorsByMovieID, id)
	if err != nil {
		return nil, fmt.Errorf("failed to query actors: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var a model.Actor
		err := rows.Scan(&a.ID, &a.Name, &a.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan actor: %w", err)
		}

		actors = append(actors, a)
	}

	return actors, nil
}

// func (r *Repository) GetActorByID(ctx context.Context, id int) {

// }

// // Admin handlers

// func (r *Repository) CreateActor(ctx context.Context) {

// }

// func (r *Repository) UpdateActor(ctx context.Context) {

// }

// func (r *Repository) DeleteActor(ctx context.Context, id int) {

// }
