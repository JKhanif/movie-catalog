package handler

import (
	"errors"
	"log"
	repo "movie_catalog/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateReview(c *gin.Context) {

}

func (h *Handler) GetReviewsByMovieID(c *gin.Context) {

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
