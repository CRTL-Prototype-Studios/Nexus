package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nexus/pkg/api/v1/models"
	"strconv"
)

func (h *Handler) CreatePhoto(c *gin.Context) {
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		FileID      uint   `json:"fileID"`
		Width       int    `json:"width"`
		Height      int    `json:"height"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	photo := models.Photo{
		Title:       input.Title,
		Description: input.Description,
		FileID:      input.FileID,
		Width:       input.Width,
		Height:      input.Height,
	}

	if err := h.DB.Create(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create photo"})
		return
	}

	c.JSON(http.StatusCreated, photo)
}

func (h *Handler) GetPhoto(c *gin.Context) {
	photoID := c.Param("id")

	var photo models.Photo
	if err := h.DB.Preload("File").First(&photo, photoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	c.JSON(http.StatusOK, photo)
}

func (h *Handler) UpdatePhoto(c *gin.Context) {
	photoID := c.Param("id")

	var photo models.Photo
	if err := h.DB.First(&photo, photoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		FileID      uint   `json:"fileID"`
		Width       int    `json:"width"`
		Height      int    `json:"height"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	photo.Title = input.Title
	photo.Description = input.Description
	photo.FileID = input.FileID
	photo.Width = input.Width
	photo.Height = input.Height

	if err := h.DB.Save(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update photo"})
		return
	}

	c.JSON(http.StatusOK, photo)
}

func (h *Handler) DeletePhoto(c *gin.Context) {
	photoID := c.Param("id")

	var photo models.Photo
	if err := h.DB.First(&photo, photoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	if err := h.DB.Delete(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}

func (h *Handler) ListPhotos(c *gin.Context) {
	var photos []models.Photo
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	offset := (page - 1) * pageSize

	if err := h.DB.Preload("File").Offset(offset).Limit(pageSize).Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch photos"})
		return
	}

	c.JSON(http.StatusOK, photos)
}
