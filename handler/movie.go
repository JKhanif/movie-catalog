package handler

import (
	"errors"
	"log"
	"movie_catalog/model"
	repo "movie_catalog/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllMovies(c *gin.Context) {

}

func (h *Handler) GetTopMovies(c *gin.Context) {
	minReviewsStr := c.DefaultQuery("min_reviews", "3")
	minReviews, err := strconv.Atoi(minReviewsStr)
	if err != nil || minReviews < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid min_reviews"})
		return
	}

	movies, err := h.repo.GetTopMovies(c, minReviews)
	if err != nil {
		log.Printf("Failed to fetch top movies: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch top movies"})
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (h *Handler) GetMovieByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Invalid movie ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	movie, err := h.repo.GetMovieByID(c, id)
	if err != nil {
		if errors.Is(err, repo.ErrMovieNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
			return
		}

		log.Printf("Failed to retrieve movie: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve movie"})
		return
	}

	c.JSON(http.StatusOK, movie)
}

// Admin handlers

func (h *Handler) CreateMovie(c *gin.Context) {
	var m model.Movie

	err := c.ShouldBindJSON(&m)
	if err != nil {
		log.Printf("Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	id, err := h.repo.CreateMovie(c, m)
	if err != nil {
		log.Printf("Failed to create movie: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create movie"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id, "message": "Movie created successfully"})
}

func (h *Handler) UpdateMovie(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Invalid movie ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	var m model.Movie
	err = c.ShouldBindJSON(&m)
	if err != nil {
		log.Printf("Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = h.repo.UpdateMovie(c, id, m)
	if err != nil {
		if errors.Is(err, repo.ErrMovieNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
			return
		}
		log.Printf("Failed to update movie: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update movie"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie updated successfully"})
}

func (h *Handler) DeleteMovie(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Invalid movie ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	err = h.repo.DeleteMovie(c, id)
	if err != nil {
		log.Printf("Failed to delete movie: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete movie"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}
