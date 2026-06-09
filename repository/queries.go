package repo

// MOVIE MOVIE MOVIE MOVIE MOVIE MOVIE MOVIE MOVIE MOVIE MOVIE MOVIE MOVIE MOVIE

const queryGetAllMovies = `SELECT m.id, m.title, m.year, m.description, m.director,
								m.created_at, m.updated_at,
								COALESCE(AVG(r.rating), 0) AS avg_rating,
								COUNT(r.id) AS review_count
							FROM movies m
							LEFT JOIN reviews r ON r.movie_id = m.id
							WHERE ($1 = '' OR EXISTS (
								SELECT 1 FROM movie_genres mg
								JOIN genres g ON g.id = mg.genre_id
								WHERE mg.movie_id = m.id AND g.name = $1
							))
							AND ($2 = 0 OR m.year = $2)
							AND ($4 = '' OR to_tsvector('russian', m.title || ' ' || COALESCE(m.description, '')) @@ plainto_tsquery('russian', $4))
							GROUP BY m.id
							HAVING ($3 = 0 OR COALESCE(AVG(r.rating), 0) >= $3)`

const queryGetTopMovies = `SELECT m.id, m.title, m.year, m.description, m.director,
								m.created_at, m.updated_at,
								COALESCE(AVG(r.rating), 0) AS avg_rating,
								COUNT(r.id) AS review_count
							FROM movies m
							LEFT JOIN reviews r ON r.movie_id = m.id
							GROUP BY m.id
							HAVING COUNT(r.id) >= $1
							ORDER BY AVG(r.rating) DESC
							LIMIT 10`

const queryGetMovieByID = `SELECT m.id, m.title, m.year, m.description, m.director,
									m.created_at, m.updated_at,
								COALESCE(AVG(r.rating), 0) AS avg_rating,
								COUNT(r.id) AS review_count
							FROM movies m
							LEFT JOIN reviews r ON r.movie_id = m.id
							WHERE m.id = $1
							GROUP BY m.id`

const queryCreateMovie = `INSERT INTO movies (title, year, description, director) VALUES ($1, $2, $3, $4) RETURNING id`

const queryDeleteMovie = `DELETE FROM movies WHERE id = $1`

// GENRE GENRE GENRE GENRE GENRE GENRE GENRE GENRE GENRE GENRE GENRE GENRE GENRE

const queryGetAllGenres = `SELECT id, name, created_at FROM genres`

const queryGetGenresByMovieID = ` SELECT g.id, g.name, g.created_at
									FROM genres g 
									JOIN movie_genres mg ON mg.genre_id = g.id 
									WHERE mg.movie_id = $1`

// const queryGetGenreByID = `SELECT id, name FROM genres WHERE id = $1`

const queryCreateGenre = `INSERT INTO genres (name) VALUES ($1)`

// const queryUpdateGenre = `UPDATE genres SET name = $1 WHERE id = $2`

const queryDeleteGenre = `DELETE FROM genres WHERE id = $1`

// REVIEW REVIEW REVIEW REVIEW REVIEW REVIEW REVIEW REVIEW REVIEW REVIEW REVIEW REVIEW REVIEW

const queryGetReviewsByMovieID = `SELECT id, movie_id, author, rating, text, created_at
									FROM reviews WHERE movie_id = $1 
									ORDER BY created_at DESC LIMIT $2 OFFSET $3`

const queryCreateReview = `INSERT INTO reviews (movie_id, author, rating, text) VALUES ($1, $2, $3, $4) RETURNING id`

// const queryUpdateReview = `UPDATE reviews SET author = $1, rating = $2, text = $3 WHERE id = $4`

const queryDeleteReview = `DELETE FROM reviews WHERE id = $1`

// ACTOR ACTOR ACTOR ACTOR ACTOR ACTOR ACTOR ACTOR ACTOR ACTOR ACTOR ACTOR ACTOR

// const queryGetAllActors = `SELECT id, name FROM actors`

const queryGetActorsByMovieID = `SELECT a.id, a.name, a.created_at FROM actors a
								JOIN movie_actors ma ON a.id = ma.actor_id
								WHERE ma.movie_id = $1`

// const queryCreateActor = `INSERT INTO actors (name) VALUES ($1)`

// const queryUpdateActor = `UPDATE actors SET name = $1 WHERE id = $2`

// const queryDeleteActor = `DELETE FROM actors WHERE id = $1`
