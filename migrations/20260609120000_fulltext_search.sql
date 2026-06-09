-- +goose Up
CREATE INDEX IF NOT EXISTS idx_movies_fts ON movies
  USING GIN (to_tsvector('russian', title || ' ' || COALESCE(description, '')));

-- +goose Down
DROP INDEX IF EXISTS idx_movies_fts;
