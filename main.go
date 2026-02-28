package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	"github.com/taufik-hdyt/go-crud/config"
	middleware "github.com/taufik-hdyt/go-crud/middlewares"
	"github.com/taufik-hdyt/go-crud/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()
	config.ConnectDatabase()
	r.Use(middleware.DBMiddleware(config.DB))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Server berjalan dan database sudah connect ✅",
		})
	})

	// routes
	api := r.Group("/api")
	routes.AuthRoutes(api)
	routes.UserRoutes(api)
	routes.CategoryRoutes(api)
	routes.ProductRoutes(api)

	r.Run(":8080")
}
