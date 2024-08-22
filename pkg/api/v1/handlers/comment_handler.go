package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"nexus/pkg/api/v1/models"
)

// AddComment handles the creation of a new comment
func (h *Handler) AddComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate that either BlogPostID or PostID is provided, but not both
	if (comment.BlogPostID == nil) || (comment.BlogPostID != nil) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Exactly one of BlogPostID or PostID must be provided"})
		return
	}

	// Check if the associated BlogPost or Post exists
	if comment.BlogPostID != nil {
		var blogPost models.BlogPost
		if err := h.DB.First(&blogPost, *comment.BlogPostID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Associated blog post not found"})
			return
		}
	}

	if err := h.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add comment"})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

// GetComments retrieves all comments for a specific blog post or post
func (h *Handler) GetComments(c *gin.Context) {
	blogPostID := c.Query("blogPostID")
	postID := c.Query("postID")

	var comments []models.Comment
	var err error

	if blogPostID != "" {
		err = h.DB.Where("blog_post_id = ?", blogPostID).Find(&comments).Error
	} else if postID != "" {
		err = h.DB.Where("post_id = ?", postID).Find(&comments).Error
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Either blogPostID or postID query parameter is required"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// DeleteComment handles the deletion of a comment
func (h *Handler) DeleteComment(c *gin.Context) {
	commentID := c.Param("id")

	var comment models.Comment
	if err := h.DB.First(&comment, commentID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comment"})
		}
		return
	}

	if err := h.DB.Delete(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

// UpdateComment handles updating an existing comment
func (h *Handler) UpdateComment(c *gin.Context) {
	commentID := c.Param("id")

	var comment models.Comment
	if err := h.DB.First(&comment, commentID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comment"})
		}
		return
	}

	var input struct {
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.Content = input.Content

	if err := h.DB.Save(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comment"})
		return
	}

	c.JSON(http.StatusOK, comment)
}
