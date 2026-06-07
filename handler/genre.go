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

func (h *Handler) GetAllGenres(c *gin.Context) {
	g, err := h.repo.GetAllGenres(c)
	if err != nil {
		log.Printf("Failed to fetch genres: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch genres"})
		return
	}

	c.JSON(http.StatusOK, g)
}

func (h *Handler) GetGenreByID(c *gin.Context) {

}

// Admin handlers

func (h *Handler) CreateGenre(c *gin.Context) {
	var g model.Genre

	err := c.ShouldBindJSON(&g)
	if err != nil {
		log.Printf("Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = h.repo.CreateGenre(c, g)
	if err != nil {
		log.Printf("Failed to create genre: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create genre"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Genre created successfully"})
}

// func (h *Handler) UpdateGenre(c *gin.Context) {

// }

func (h *Handler) DeleteGenre(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Invalid genre ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid genre ID"})
		return
	}

	err = h.repo.DeleteGenre(c, id)
	if err != nil {
		if errors.Is(err, repo.ErrGenreNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Genre not found"})
			return
		}

		log.Printf("Failed to delete genre: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete genre"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Genre deleted successfully"})
}
