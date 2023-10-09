package handlers

import (
	"UrlShortenerV1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllURLs retrieves all shortened URLs
func GetAllURLs(c *gin.Context) {
	// This is a mock data
	urls := []models.URL{
		{ID: "1", OriginalURL: "https://example.com", ShortURL: "https://short.url/e1"},
	}
	c.JSON(http.StatusOK, urls)
}
