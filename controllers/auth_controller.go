package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/taufik-hdyt/go-crud/config"
	"github.com/taufik-hdyt/go-crud/helpers"
	"github.com/taufik-hdyt/go-crud/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func Register(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
		Avatar   string `json:"avatar"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cek email duplikat
	var existing models.User
	if err := config.DB.Where("email = ?", input.Email).First(&existing).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email sudah digunakan"})
		return
	}

	// Hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Avatar:   input.Avatar,
	}

	config.DB.Create(&user)

	c.JSON(http.StatusCreated, gin.H{"message": "Registrasi berhasil", "user": user})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		helpers.Error(c, http.StatusUnauthorized, "Email tidak ditemukan")
		return
	}

	// Verifikasi password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		helpers.Error(c, http.StatusUnauthorized, "Password salah")
		return
	}

	// Buat JWT token
	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // berlaku 24 jam
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, "Gagal membuat token")
		return
	}

	helpers.Success(c, "Login berhasil","token", tokenString)
}

func ProfileMe(c *gin.Context) {
	userID, exists := c.Get("user_id")
	fmt.Println(userID)
	if !exists {
		helpers.Error(c, http.StatusNotFound, "User tidak ditemukan di context")
		return
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		helpers.Error(c, http.StatusNotFound, "User tidak ditemukan")
		return
	}
		helpers.Success(c, "Login berhasil","user", user)
}
