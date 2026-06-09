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
	genre := c.Query("genre")
	search := c.Query("search")

	yearStr := c.Query("year")
	year := 0
	if yearStr != "" {
		var err error
		year, err = strconv.Atoi(yearStr)
		if err != nil || year < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid year"})
			return
		}
	}

	minRatingStr := c.Query("min_rating")
	var minRating float64
	if minRatingStr != "" {
		var err error
		minRating, err = strconv.ParseFloat(minRatingStr, 64)
		if err != nil || minRating < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid min_rating"})
			return
		}
	}

	sort := c.DefaultQuery("sort", "created_at")
	order := c.DefaultQuery("order", "desc")

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil || offset < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
		return
	}

	movies, err := h.repo.GetAllMovies(c, genre, year, minRating, search, sort, order, limit, offset)
	if err != nil {
		log.Printf("Failed to fetch movies: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies"})
		return
	}

	c.JSON(http.StatusOK, movies)
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

	var req model.UpdateMovieRequest
	err = c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = h.repo.UpdateMovie(c, id, req)
	if err != nil {
		if errors.Is(err, repo.ErrMovieNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
			return
		}
		if errors.Is(err, repo.ErrNoFieldsToUpdate) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No fields to update"})
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
		if errors.Is(err, repo.ErrMovieNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
			return
		}
		log.Printf("Failed to delete movie: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete movie"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}
