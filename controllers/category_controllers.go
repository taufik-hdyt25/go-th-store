package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/taufik-hdyt/go-crud/models"
	"gorm.io/gorm"
)

func CreateCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input models.Category

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := models.Category{
		Name:      input.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db.Create(&category)
	c.JSON(http.StatusCreated, gin.H{"data": category})
}

func GetCategories(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var categories []models.Category
	db.Find(&categories)
	c.JSON(http.StatusOK, gin.H{"data": categories})
}

func GetCategoryByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var category models.Category
	id := c.Param("id")

	if err := db.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

func UpdateCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var category models.Category
	id := c.Param("id")

	if err := db.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category.Name = input.Name
	category.UpdatedAt = time.Now()

	db.Save(&category)
	c.JSON(http.StatusOK, gin.H{"data": category})
}

func DeleteCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var category models.Category
	id := c.Param("id")

	if err := db.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	db.Delete(&category)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
