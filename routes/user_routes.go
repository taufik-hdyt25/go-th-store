package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/taufik-hdyt/go-crud/controllers"
)

func UserRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{

		users.POST("/", controllers.Register)
		users.GET("/", controllers.GetUsers)
		users.PUT("/:id", controllers.UpdateUser)
		users.DELETE("/:id", controllers.DeleteUser)
	}
}
