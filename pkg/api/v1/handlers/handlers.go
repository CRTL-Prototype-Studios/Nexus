package handlers

import (
	"github.com/minio/minio-go/v7"
	"net/http"
	models2 "nexus/pkg/api/v1/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB          *gorm.DB
	MinioClient *minio.Client
}

func (h *Handler) CreateBlogPost(c *gin.Context) {
	var blogPost models2.BlogPost
	if err := c.ShouldBindJSON(&blogPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Create(&blogPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create blog post"})
		return
	}

	c.JSON(http.StatusCreated, blogPost)
}

func (h *Handler) GetBlogPosts(c *gin.Context) {
	var blogPosts []models2.BlogPost
	if err := h.DB.Find(&blogPosts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch blog posts"})
		return
	}

	c.JSON(http.StatusOK, blogPosts)
}

// Add similar functions for Photo and Post CRUD operations

func (h *Handler) AddComment(c *gin.Context) {
	var comment models2.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add comment"})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

// Add a function for file upload to MinIO
func (h *Handler) UploadFile(c *gin.Context) {
	// Implementation for file upload to MinIO
	// This will involve handling multipart form data and using the MinIO client to upload the file
}
