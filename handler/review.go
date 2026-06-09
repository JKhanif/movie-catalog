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

func (h *Handler) CreateReview(c *gin.Context) {
	movieIDStr := c.Param("id")
	movieID, err := strconv.Atoi(movieIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	var input model.CreateReviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Валидация
	if input.Author == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Author is required"})
		return
	}
	if input.Rating < 1 || input.Rating > 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Rating must be between 1 and 10"})
		return
	}

	id, err := h.repo.CreateReview(c, movieID, input.Author, input.Rating, input.Text)
	if err != nil {
		log.Printf("Failed to create review: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create review"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *Handler) GetReviewsByMovieID(c *gin.Context) {
	movieIDStr := c.Param("id")
	movieID, err := strconv.Atoi(movieIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

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

	reviews, err := h.repo.GetReviewsByMovieID(c, movieID, limit, offset)
	if err != nil {
		log.Printf("Failed to fetch reviews: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reviews"})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

func (h *Handler) DeleteReview(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Invalid review ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		return
	}

	err = h.repo.DeleteReview(c, id)
	if err != nil {
		if errors.Is(err, repo.ErrReviewNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
			return
		}

		log.Printf("Failed to delete review: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete review"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}
