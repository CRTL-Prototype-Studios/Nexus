package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"nexus/pkg/api/v1/models"
)

// CreateAlbum creates a new album
func (h *Handler) CreateAlbum(c *gin.Context) {
	var album models.Album
	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Create(&album).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create album"})
		return
	}

	c.JSON(http.StatusCreated, album)
}

// GetAlbum retrieves a specific album
func (h *Handler) GetAlbum(c *gin.Context) {
	id := c.Param("id")

	var album models.Album
	if err := h.DB.Preload("Photos").First(&album, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
		return
	}

	c.JSON(http.StatusOK, album)
}

// ListAlbums retrieves all albums
func (h *Handler) ListAlbums(c *gin.Context) {
	var albums []models.Album
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	offset := (page - 1) * pageSize

	if err := h.DB.Offset(offset).Limit(pageSize).Find(&albums).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch albums"})
		return
	}

	c.JSON(http.StatusOK, albums)
}

// UpdateAlbum updates an existing album
func (h *Handler) UpdateAlbum(c *gin.Context) {
	id := c.Param("id")

	var album models.Album
	if err := h.DB.First(&album, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
		return
	}

	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Save(&album).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update album"})
		return
	}

	c.JSON(http.StatusOK, album)
}

// DeleteAlbum deletes an album
func (h *Handler) DeleteAlbum(c *gin.Context) {
	id := c.Param("id")

	if err := h.DB.Delete(&models.Album{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete album"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Album deleted successfully"})
}

// AddPhotoToAlbum adds a photo to an album
func (h *Handler) AddPhotoToAlbum(c *gin.Context) {
	albumID := c.Param("id")
	var input struct {
		PhotoID uint `json:"photoID" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var album models.Album
	if err := h.DB.First(&album, albumID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
		return
	}

	var photo models.Photo
	if err := h.DB.First(&photo, input.PhotoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	if err := h.DB.Model(&album).Association("Photos").Append(&photo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add photo to album"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo added to album successfully"})
}

// RemovePhotoFromAlbum removes a photo from an album
func (h *Handler) RemovePhotoFromAlbum(c *gin.Context) {
	albumID := c.Param("id")
	photoID := c.Param("photoID")

	var album models.Album
	if err := h.DB.First(&album, albumID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
		return
	}

	var photo models.Photo
	if err := h.DB.First(&photo, photoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	if err := h.DB.Model(&album).Association("Photos").Delete(&photo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove photo from album"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo removed from album successfully"})
}
