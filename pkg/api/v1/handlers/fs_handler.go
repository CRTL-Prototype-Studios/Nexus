package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
	"net/http"
	"nexus/pkg/api/v1/models"
	"path/filepath"
	"strings"
	"time"
)

func (h *Handler) GetFile(c *gin.Context) {
	fileID := c.Param("id")

	var file models.File
	if err := h.DB.First(&file, fileID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	c.JSON(http.StatusOK, file)
}

func (h *Handler) UpdateFile(c *gin.Context) {
	fileID := c.Param("id")

	var file models.File
	if err := h.DB.First(&file, fileID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	var input struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file.Name = input.Name

	if err := h.DB.Save(&file).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update file"})
		return
	}

	c.JSON(http.StatusOK, file)
}

func (h *Handler) DeleteFile(c *gin.Context) {
	fileID := c.Param("id")

	var file models.File
	if err := h.DB.First(&file, fileID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// Delete the file from MinIO
	err := h.MinioClient.RemoveObject(c, h.BucketName, filepath.Base(file.URL), minio.RemoveObjectOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file from storage"})
		return
	}

	// Delete the file record from the database
	if err := h.DB.Delete(&file).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}

func (h *Handler) UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file from unable to close file writing"})
		}
	}(file)

	path := c.DefaultPostForm("path", "")
	isDirectory := c.DefaultPostForm("isDirectory", "false") == "true"

	// Generate a unique filename
	filename := uuid.New().String() + filepath.Ext(header.Filename)
	fullPath := filepath.Join(path, filename)

	if !isDirectory {
		// Upload the file to MinIO
		_, err = h.MinioClient.PutObject(c, h.BucketName, fullPath, file, header.Size, minio.PutObjectOptions{ContentType: header.Header.Get("Content-Type")})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
			return
		}
	}

	// Create a new File record in the database
	fileRecord := models.File{
		Name:        header.Filename,
		Path:        path,
		Size:        header.Size,
		URL:         fmt.Sprintf("/%s/%s", h.BucketName, fullPath),
		ContentType: header.Header.Get("Content-Type"),
		UploadedAt:  time.Now(),
		IsDirectory: isDirectory,
	}

	if err := h.DB.Create(&fileRecord).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file record"})
		return
	}

	c.JSON(http.StatusOK, fileRecord)
}

func (h *Handler) ListFiles(c *gin.Context) {
	path := c.DefaultQuery("path", "/")
	var files []models.File

	if err := h.DB.Where("path = ?", path).Find(&files).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch files"})
		return
	}

	c.JSON(http.StatusOK, files)
}

func (h *Handler) GetFileByPath(c *gin.Context) {
	path := c.Param("path")
	path = strings.TrimPrefix(path, "/")

	var files []models.File
	if err := h.DB.Where("path = ?", path).Or("path LIKE ?", path+"/%").Find(&files).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch files"})
		return
	}

	// If it's a single file, return it directly
	if len(files) == 1 && !files[0].IsDirectory {
		c.JSON(http.StatusOK, files[0])
		return
	}

	// Otherwise, return the list of files/directories
	c.JSON(http.StatusOK, files)
}

// Other methods (GetFile, UpdateFile, DeleteFile) remain the same
// Remember to update them to handle the new 'Path' field if necessary

func (h *Handler) CreateDirectory(c *gin.Context) {
	var input struct {
		Name string `json:"name" binding:"required"`
		Path string `json:"path" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fullPath := filepath.Join(input.Path, input.Name)

	directory := models.File{
		Name:        input.Name,
		Path:        input.Path,
		Size:        0,
		URL:         fmt.Sprintf("/%s/%s", h.BucketName, fullPath),
		ContentType: "directory",
		UploadedAt:  time.Now(),
		IsDirectory: true,
	}

	if err := h.DB.Create(&directory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create directory"})
		return
	}

	c.JSON(http.StatusCreated, directory)
}
