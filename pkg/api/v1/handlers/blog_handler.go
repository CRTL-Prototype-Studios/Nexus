package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"nexus/pkg/api/v1/models"
)

func (h *Handler) CreateBlogPost(c *gin.Context) {
	var blogPost models.BlogPost
	if err := c.BindJSON(&blogPost); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if blogPost.Comments == nil {
		blogPost.Comments = []models.Comment{}
	}

	// Handle the case where no cover is provided
	if blogPost.CoverID == nil {
		blogPost.Cover = nil
	} else {
		// If a CoverID is provided, fetch the corresponding Photo
		var cover models.Photo
		if err := h.DB.First(&cover, *blogPost.CoverID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Provided CoverID does not exist"})
			} else {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cover photo"})
			}
			return
		}
		blogPost.Cover = &cover
	}

	if err := h.DB.Create(&blogPost).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create blog post"})
		return
	}

	c.IndentedJSON(http.StatusCreated, blogPost)
}

func (h *Handler) GetBlogPosts(c *gin.Context) {
	var blogPosts []models.BlogPost
	if err := h.DB.Find(&blogPosts).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, blogPosts)
}

func (h *Handler) GetBlogPost(c *gin.Context) {
	var blogPost models.BlogPost
	if err := h.DB.First(&blogPost, c.Param("id")).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, blogPost)
}

func (h *Handler) UpdateBlogPost(c *gin.Context) {
	// Get the blog post ID from the URL parameter
	id := c.Param("id")

	// Find the existing blog post
	var existingPost models.BlogPost
	if err := h.DB.First(&existingPost, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Blog post not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch blog post"})
		}
		return
	}

	// Bind the JSON input to a struct
	var input struct {
		Title   *string `json:"title"`
		Content *string `json:"content"`
		CoverID *uint   `json:"coverID"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields if they're provided
	if input.Title != nil {
		existingPost.Title = *input.Title
	}
	if input.Content != nil {
		existingPost.Content = *input.Content
	}

	// Handle cover update
	if input.CoverID != nil {
		if *input.CoverID == 0 {
			// Remove cover
			existingPost.CoverID = nil
			existingPost.Cover = nil
		} else {
			// Update cover
			var cover models.Photo
			if err := h.DB.First(&cover, *input.CoverID).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Provided CoverID does not exist"})
				} else {
					c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cover photo"})
				}
				return
			}
			existingPost.CoverID = input.CoverID
			existingPost.Cover = &cover
		}
	}

	// Save the updated blog post
	if err := h.DB.Save(&existingPost).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog post"})
		return
	}

	c.IndentedJSON(http.StatusOK, existingPost)
}
