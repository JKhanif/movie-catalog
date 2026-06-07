package repo

// MOVIE MOVIE MOVIE MOVIE MOVIE MOVIE MOVIE MOVIE MOVIE MOVIE MOVIE MOVIE MOVIE

const queryGetAllMovies = `SELECT id, title, year, description, director, created_at, updated_at FROM movies`

const queryGetTopMovies = `SELECT id, title, year, description, director, created_at, updated_at
							FROM movies m
							LEFT JOIN reviews r ON m.id = r.movie_id
							GROUP BY m.id
							ORDER BY AVG(r.rating) DESC
							LIMIT 10`

const queryGetMovieByID = `SELECT id, title, year, description, director, created_at, updated_at 
    						FROM movies WHERE id = $1`

const queryCreateMovie = `INSERT INTO movies (title, year, description, director) VALUES ($1, $2, $3, $4) RETURNING id`

const queryUpdateMovie = `UPDATE movies SET title = $1, year = $2, description = $3, director = $4 WHERE id = $5`

const queryDeleteMovie = `DELETE FROM movies WHERE id = $1`

// GENRE GENRE GENRE GENRE GENRE GENRE GENRE GENRE GENRE GENRE GENRE GENRE GENRE

const queryGetAllGenres = `SELECT id, name FROM genres`

// const queryGetGenreByID = `SELECT id, name FROM genres WHERE id = $1`

const queryCreateGenre = `INSERT INTO genres (name) VALUES ($1)`

// const queryUpdateGenre = `UPDATE genres SET name = $1 WHERE id = $2`

const queryDeleteGenre = `DELETE FROM genres WHERE id = $1`

// REVIEW REVIEW REVIEW REVIEW REVIEW REVIEW REVIEW REVIEW REVIEW REVIEW REVIEW REVIEW REVIEW

const queryGetReviewsByMovieID = `SELECT id, movie_id, author, rating, text, created_at
									FROM reviews WHERE movie_id = $1`

const queryCreateReview = `INSERT INTO reviews (movie_id, author, rating, text) VALUES ($1, $2, $3, $4)`

// const queryUpdateReview = `UPDATE reviews SET author = $1, rating = $2, text = $3 WHERE id = $4`

const queryDeleteReview = `DELETE FROM reviews WHERE id = $1`

// ACTOR ACTOR ACTOR ACTOR ACTOR ACTOR ACTOR ACTOR ACTOR ACTOR ACTOR ACTOR ACTOR

// const queryGetAllActors = `SELECT id, name FROM actors`

const queryGetActorsByMovieID = `SELECT name FROM actors a
								JOIN movie_actors ma ON a.id = ma.actor_id
								WHERE ma.movie_id = $1`

// const queryCreateActor = `INSERT INTO actors (name) VALUES ($1)`

// const queryUpdateActor = `UPDATE actors SET name = $1 WHERE id = $2`

// const queryDeleteActor = `DELETE FROM actors WHERE id = $1`
