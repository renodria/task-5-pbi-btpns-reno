package router

import (
	"vix-btpns/controllers"
	"vix-btpns/middlewares"

	"github.com/gin-gonic/gin"
)

func PhotosRouter(r *gin.Engine) {
	photosGroup := r.Group("/photos")
	photosGroup.Use(middlewares.Authentication())

	photosGroup.GET("/", controllers.ShowPhotos)
	photosGroup.GET("/:photoId", controllers.ShowPhotoId)
	photosGroup.POST("/", controllers.CreatePhoto)
	photosGroup.PUT("/:photoId", controllers.UpdatePhoto)
	photosGroup.DELETE("/:photoId", controllers.DeletePhoto)
}
