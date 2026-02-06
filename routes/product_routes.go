package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/taufik-hdyt/go-crud/controllers"
	middleware "github.com/taufik-hdyt/go-crud/middlewares"
)

func ProductRoutes(r *gin.Engine) {
	product := r.Group("/category")
	product.Use(middleware.AuthMiddleware())
	{
		product.POST("/products", controllers.CreateProduct)
		product.GET("/products", controllers.GetProducts)
		product.GET("/products/:id", controllers.GetProductByID)
		product.PUT("/products/:id", controllers.UpdateProduct)
		product.DELETE("/products/:id", controllers.DeleteProduct)
	}
}
