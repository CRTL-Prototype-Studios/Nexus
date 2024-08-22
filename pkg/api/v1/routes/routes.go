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
		blog := v1.Group("/blog")
		{
			blog.POST("/", h.CreateBlogPost)
			blog.GET("/", h.GetBlogPosts)
			blog.GET("/:id", h.GetBlogPost)
			blog.PUT("/:id", h.UpdateBlogPost)
		}

		files := v1.Group("/files")
		{
			files.POST("/", h.UploadFile)
			files.GET("/", h.ListFiles)
			files.GET("/*path", h.GetFileByPath)
			files.PUT("/:id", h.UpdateFile)
			files.DELETE("/:id", h.DeleteFile)
		}

		photos := v1.Group("/photos")
		{
			// Photo routes
			photos.POST("/", h.CreatePhoto)
			photos.GET("/:id", h.GetPhoto)
			photos.PUT("/:id", h.UpdatePhoto)
			photos.DELETE("/:id", h.DeletePhoto)
			photos.GET("/", h.ListPhotos)
		}

		albums := router.Group("/albums")
		{
			albums.POST("/", h.CreateAlbum)
			albums.GET("/", h.ListAlbums)
			albums.GET("/:id", h.GetAlbum)
			albums.PUT("/:id", h.UpdateAlbum)
			albums.DELETE("/:id", h.DeleteAlbum)
			albums.POST("/:id/photos", h.AddPhotoToAlbum)
			albums.DELETE("/:id/photos/:photoID", h.RemovePhotoFromAlbum)
		}

		// Directory routes
		v1.POST("/directories", h.CreateDirectory)

		// Comment routes
		comments := v1.Group("/comments")
		{
			comments.POST("/", h.AddComment)
			comments.GET("/", h.GetComments)
			comments.DELETE("/:id", h.DeleteComment)
			comments.PUT("/:id", h.UpdateComment)
		}

		// File upload route
		// v1.POST("/upload", h.UploadFile)
	}
}
