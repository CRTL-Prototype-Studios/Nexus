package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
	"nexus/pkg/api/v1/handlers"
	"nexus/pkg/api/v1/middleware"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB, minioClient *minio.Client) {
	h := &handlers.Handler{DB: db, MinioClient: minioClient}

	// Public routes
	router.POST("/signup", h.SignUp)
	router.POST("/signin", h.SignIn)

	// Apply rate limiting to all routes
	router.Use(middleware.RateLimitMiddleware())

	v1 := router.Group("/api/v1")
	{
		// Blog routes
		blog := v1.Group("/blog")
		{
			blog.GET("/", h.GetBlogPosts)
			blog.GET("/:id", h.GetBlogPost)

			// Protected routes
			authBlog := blog.Group("/")
			authBlog.Use(middleware.AuthMiddleware())
			{
				authBlog.POST("/", middleware.RBACMiddleware("create_post"), h.CreateBlogPost)
				authBlog.PUT("/:id", middleware.RBACMiddleware("edit_post"), h.UpdateBlogPost)
			}
		}

		// File routes
		files := v1.Group("/files")
		{
			files.GET("/", h.ListFiles)

			// Protected routes
			authFiles := files.Group("/")
			authFiles.Use(middleware.AuthMiddleware())
			{
				authFiles.POST("/", middleware.RBACMiddleware("upload_file"), h.UploadFile)
				authFiles.PUT("/:id", middleware.RBACMiddleware("edit_file"), h.UpdateFile)
				authFiles.DELETE("/:id", middleware.RBACMiddleware("delete_file"), h.DeleteFile)
			}

			// This catch-all route should be last
			files.GET("/dir/*path", h.GetFileByPath)
		}

		// Photo routes
		photos := v1.Group("/photos")
		{
			photos.GET("/:id", h.GetPhoto)
			photos.GET("/", h.ListPhotos)

			// Protected routes
			authPhotos := photos.Group("/")
			authPhotos.Use(middleware.AuthMiddleware())
			{
				authPhotos.POST("/", middleware.RBACMiddleware("create_photo"), h.CreatePhoto)
				authPhotos.PUT("/:id", middleware.RBACMiddleware("edit_photo"), h.UpdatePhoto)
				authPhotos.DELETE("/:id", middleware.RBACMiddleware("delete_photo"), h.DeletePhoto)
			}
		}

		// Album routes
		albums := v1.Group("/albums")
		{
			albums.GET("/", h.ListAlbums)
			albums.GET("/:id", h.GetAlbum)

			// Protected routes
			authAlbums := albums.Group("/")
			authAlbums.Use(middleware.AuthMiddleware())
			{
				authAlbums.POST("/", middleware.RBACMiddleware("create_album"), h.CreateAlbum)
				authAlbums.PUT("/:id", middleware.RBACMiddleware("edit_album"), h.UpdateAlbum)
				authAlbums.DELETE("/:id", middleware.RBACMiddleware("delete_album"), h.DeleteAlbum)
				authAlbums.POST("/:id/photos", middleware.RBACMiddleware("add_photo_to_album"), h.AddPhotoToAlbum)
				authAlbums.DELETE("/:id/photos/:photoID", middleware.RBACMiddleware("remove_photo_from_album"), h.RemovePhotoFromAlbum)
			}
		}

		// Directory routes
		directories := v1.Group("/directories")
		directories.Use(middleware.AuthMiddleware())
		{
			directories.POST("/", middleware.RBACMiddleware("create_directory"), h.CreateDirectory)
		}

		// Comment routes
		comments := v1.Group("/comments")
		{
			comments.GET("/", h.GetComments)

			// Protected routes
			authComments := comments.Group("/")
			authComments.Use(middleware.AuthMiddleware())
			{
				authComments.POST("/", middleware.RBACMiddleware("add_comment"), h.AddComment)
				authComments.DELETE("/:id", middleware.RBACMiddleware("delete_comment"), h.DeleteComment)
				authComments.PUT("/:id", middleware.RBACMiddleware("edit_comment"), h.UpdateComment)
			}
		}
	}
}
