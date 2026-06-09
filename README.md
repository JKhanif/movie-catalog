Миграции 
goose -dir migrations postgres "postgres://postgres:postgres@localhost:5432/movie_catalog?sslmode=disable" up

Залить seed
docker cp migrations/scripts/seed.sql movie_catalog:/seed.sql
docker exec movie_catalog psql -U postgres -d movie_catalog -f /seed.sql

Проверить, что данные есть:
docker exec movie_catalog psql -U postgres -d movie_catalog -c "SELECT title, year FROM movies;"

## API Examples

### Public

```bash
# List all movies (filters, sort, pagination)
curl "http://localhost:8080/movies"
curl "http://localhost:8080/movies?genre=Action&year=2024&min_rating=7&sort=rating&order=desc&limit=5&offset=0"

# Top movies
curl "http://localhost:8080/movies/top?min_reviews=3"

# Movie by ID (with genres, actors, stats)
curl "http://localhost:8080/movies/1"

# Reviews by movie
curl "http://localhost:8080/movies/1/reviews"

# Create review
curl -X POST "http://localhost:8080/movies/1/reviews" \
  -H "Content-Type: application/json" \
  -d '{"author":"Alice","rating":8,"text":"Great movie!"}'

# All genres
curl "http://localhost:8080/genres"
```

### Admin (require `Authorization: Bearer <token>`)

```bash
# Create movie
curl -X POST "http://localhost:8080/movies" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer admin" \
  -d '{"title":"New Movie","year":2025,"description":"...","director":"..."}'

# Update movie
curl -X PATCH "http://localhost:8080/movies/1" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer admin" \
  -d '{"title":"Updated Title"}'

# Delete movie
curl -X DELETE "http://localhost:8080/movies/1" \
  -H "Authorization: Bearer admin"

# Create genre
curl -X POST "http://localhost:8080/genres" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer admin" \
  -d '{"name":"New Genre"}'

# Delete genre
curl -X DELETE "http://localhost:8080/genres/1" \
  -H "Authorization: Bearer admin"

# Delete review
curl -X DELETE "http://localhost:8080/reviews/1" \
  -H "Authorization: Bearer admin"
```