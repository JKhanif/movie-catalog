Миграции 
goose -dir migrations postgres "postgres://postgres:postgres@localhost:5432/movie_catalog?sslmode=disable" up

Залить seed
docker cp migrations/scripts/seed.sql movie_catalog:/seed.sql
docker exec movie_catalog psql -U postgres -d movie_catalog -f /seed.sql

Проверить, что данные есть:
docker exec movie_catalog psql -U postgres -d movie_catalog -c "SELECT title, year FROM movies;"

