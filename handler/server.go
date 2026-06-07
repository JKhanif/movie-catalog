package handler

import (
	repo "movie_catalog/repository"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo *repo.Repository
}

func New(repo *repo.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) Run() {
	r := gin.Default()

	r.GET("/movies", h.GetAllMovies)
	r.GET("/movies/top", h.GetTopMovies)
	r.GET("/movies/:id", h.GetMovieByID)
	r.GET("/movies/:id/reviews", h.GetReviewsByMovieID)
	r.POST("/movies/:id/reviews", h.CreateReview)
	r.GET("/genres", h.GetAllGenres)

	admin := r.Group("/", h.adminMiddleware())
	admin.POST("/movies", h.CreateMovie)
	admin.PATCH("/movies/:id", h.UpdateMovie)
	admin.DELETE("/movies/:id", h.DeleteMovie)
	admin.POST("/genres", h.CreateGenre)
	admin.DELETE("/genres/:id", h.DeleteGenre)

	r.Run()
}
