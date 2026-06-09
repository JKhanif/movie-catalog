## Запуск

```bash
# 1. Настройка БД (PostgreSQL)
#    Через Docker:
docker run -d --name movie_catalog -p 5432:5432 \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=movie_catalog \
  postgres:16

#    Или через свой локальный PostgreSQL, создав БД movie_catalog

# 2. Применение миграций (нужен goose: go install github.com/pressly/goose/v3/cmd/goose@latest)
goose -dir migrations postgres "postgres://postgres:postgres@localhost:5432/movie_catalog?sslmode=disable" up

# 3. Залить seed-данные
psql -U postgres -d movie_catalog -f migrations/scripts/seed.sql

# 4. Запустить сервер
go run .
```

## API Examples

### Public

```bash
# List all movies (filters, sort, pagination, full-text search)
curl "http://localhost:8080/movies"
curl "http://localhost:8080/movies?genre=Фантастика&year=2014&min_rating=7&sort=rating&order=desc&limit=5&offset=0"
curl "http://localhost:8080/movies?search=интерстеллар"

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

Токен по умолчанию: `my-super-secret-token` (задаётся в .env через `ADMIN_TOKEN`)

```bash
# Create movie
curl -X POST "http://localhost:8080/movies" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer my-super-secret-token" \
  -d '{"title":"New Movie","year":2025,"description":"...","director":"..."}'

# Update movie (частичное обновление)
curl -X PATCH "http://localhost:8080/movies/1" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer my-super-secret-token" \
  -d '{"title":"Updated Title"}'

# Delete movie
curl -X DELETE "http://localhost:8080/movies/1" \
  -H "Authorization: Bearer my-super-secret-token"

# Create genre
curl -X POST "http://localhost:8080/genres" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer my-super-secret-token" \
  -d '{"name":"New Genre"}'

# Delete genre
curl -X DELETE "http://localhost:8080/genres/1" \
  -H "Authorization: Bearer my-super-secret-token"

# Delete review
curl -X DELETE "http://localhost:8080/reviews/1" \
  -H "Authorization: Bearer my-super-secret-token"
```
