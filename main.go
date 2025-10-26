package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/taufik-hdyt/go-crud/config"
	middleware "github.com/taufik-hdyt/go-crud/middlewares"
	"github.com/taufik-hdyt/go-crud/routes"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()
	r.Use(middleware.DBMiddleware(config.DB))

	routes.UserRoutes(r)
	routes.AuthRoutes(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Server berjalan dan database sudah connect ✅",
		})
	})

	r.Run(":8080")
}
