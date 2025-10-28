package middleware

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak ditemukan"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		secret := os.Getenv("JWT_SECRET")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			c.Abort()
			return
		}

		// Ambil claims dari token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Claims tidak valid"})
			c.Abort()
			return
		}

		// Ambil user_id dari claims (bisa float64 atau string)
		var userID uint
		switch v := claims["user_id"].(type) {
		case float64:
			userID = uint(v)
		case string:
			id, _ := strconv.ParseUint(v, 10, 64)
			userID = uint(id)
		default:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id tidak ditemukan di token"})
			c.Abort()
			return
		}

		// Set ke context supaya bisa diakses di controller
		c.Set("user_id", userID)

		c.Next()
	}
}
func DBMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
