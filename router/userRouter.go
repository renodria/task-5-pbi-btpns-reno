package router

import (
	"vix-btpns/controllers"
	"vix-btpns/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	userGroup := r.Group("/users")

	userGroup.POST("/register", controllers.Register)
	userGroup.POST("/login", controllers.Login)

	userGroup.Use(middlewares.Authentication())

	userGroup.PUT("/:userId", controllers.UpdateUser)
	userGroup.DELETE("/:userId", controllers.DeleteUser)
}
