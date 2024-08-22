package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
	"nexus/pkg/api/v1/handlers"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB, minioClient *minio.Client) {
	h := &handlers.Handler{DB: db, MinioClient: minioClient}

	v1 := router.Group("/api/v1")
	{
		// Blog routes
		v1.POST("/blog", h.CreateBlogPost)
		v1.GET("/blog", h.GetBlogPosts)
		v1.GET("/blog/:id", h.GetBlogPost)
		v1.PUT("/blog/:id", h.UpdateBlogPost)

		// Photo routes
		// v1.POST("/photo", h.CreatePhoto)
		// v1.GET("/photo", h.GetPhotos)

		// Post routes
		// v1.POST("/post", h.CreatePost)
		// v1.GET("/post", h.GetPosts)

		// Comment routes
		v1.POST("/comments", h.AddComment)
		v1.GET("/comments", h.GetComments)
		v1.DELETE("/comments/:id", h.DeleteComment)
		v1.PUT("/comments/:id", h.UpdateComment)

		// File upload route
		v1.POST("/upload", h.UploadFile)
	}
}
