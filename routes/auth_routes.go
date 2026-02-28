package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/taufik-hdyt/go-crud/controllers"
	middleware "github.com/taufik-hdyt/go-crud/middlewares"
)

func AuthRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	auth.POST("/login", controllers.Login)
	auth.POST("/register", controllers.Register)

	protected := r.Group("/auth")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/me", controllers.ProfileMe)
}
