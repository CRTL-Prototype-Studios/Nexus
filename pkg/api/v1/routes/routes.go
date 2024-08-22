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

		v1.POST("/files", h.UploadFile)
		v1.GET("/files", h.ListFiles)
		v1.GET("/files/*path", h.GetFileByPath)
		v1.PUT("/files/:id", h.UpdateFile)
		v1.DELETE("/files/:id", h.DeleteFile)

		// Photo routes
		v1.POST("/photos", h.CreatePhoto)
		v1.GET("/photos/:id", h.GetPhoto)
		v1.PUT("/photos/:id", h.UpdatePhoto)
		v1.DELETE("/photos/:id", h.DeletePhoto)
		v1.GET("/photos", h.ListPhotos)

		// Directory routes
		v1.POST("/directories", h.CreateDirectory)

		// Comment routes
		v1.POST("/comments", h.AddComment)
		v1.GET("/comments", h.GetComments)
		v1.DELETE("/comments/:id", h.DeleteComment)
		v1.PUT("/comments/:id", h.UpdateComment)

		// File upload route
		v1.POST("/upload", h.UploadFile)
	}
}
