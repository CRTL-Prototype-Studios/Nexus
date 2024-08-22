package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/yourusername/mywebsite-backend/api/handlers"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB, minioClient *minio.Client) {
	h := &handlers.Handler{DB: db, MinioClient: minioClient}

	v1 := router.Group("/api/v1")
	{
		// Blog routes
		v1.POST("/blog", h.CreateBlogPost)
		v1.GET("/blog", h.GetBlogPosts)

		// Photo routes
		// v1.POST("/photo", h.CreatePhoto)
		// v1.GET("/photo", h.GetPhotos)

		// Post routes
		// v1.POST("/post", h.CreatePost)
		// v1.GET("/post", h.GetPosts)

		// Comment routes
		v1.POST("/comment", h.AddComment)

		// File upload route
		v1.POST("/upload", h.UploadFile)
	}
}
