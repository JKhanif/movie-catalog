package main

import (
	"context"
	"fmt"
	"log"
	"movie_catalog/handler"
	repo "movie_catalog/repository"
	"os"
	"os/signal"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found, using environment variables: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer pool.Close()

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}

	repo := repo.New(pool)
	h := handler.New(repo)

	h.Run()

	// ждём Ctrl+C
	<-ctx.Done()
	fmt.Println("Shutting down...")
}
