package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/taufik-hdyt/go-crud/controllers"
	middleware "github.com/taufik-hdyt/go-crud/middlewares"
)

func CategoryRoutes(r *gin.RouterGroup) {
	category := r.Group("/category")
	category.Use(middleware.AuthMiddleware())
	{
		r.POST("/categories", controllers.CreateCategory)
		r.GET("/categories", controllers.GetCategories)
		r.GET("/categories/:id", controllers.GetCategoryByID)
		r.PUT("/categories/:id", controllers.UpdateCategory)
		r.DELETE("/categories/:id", controllers.DeleteCategory)
	}
}
