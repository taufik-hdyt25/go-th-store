package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/taufik-hdyt/go-crud/controllers"
	middleware "github.com/taufik-hdyt/go-crud/middlewares"
)

func UserRoutes(r *gin.RouterGroup) {
	users := r.Group("/users")
	users.Use(middleware.AuthMiddleware())
	{
		users.GET("/", controllers.GetUsers)
		users.GET("/:id", controllers.GetUser)
		users.PUT("/:id", controllers.UpdateUser)
		users.DELETE("/:id", controllers.DeleteUser)
	}
}
