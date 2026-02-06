package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/taufik-hdyt/go-crud/models"
	"gorm.io/gorm"
)

// CREATE
func CreateProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input models.Product

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := models.Product{
		CategoryID: input.CategoryID,
		Name:       input.Name,
		Stok:       input.Stok,
		Harga:      input.Harga,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := db.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": product})
}

// READ ALL
func GetProducts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var products []models.Product

	if err := db.Preload("Category").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// READ ONE
func GetProductByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var product models.Product
	id := c.Param("id")

	if err := db.Preload("Category").First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// UPDATE
func UpdateProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var product models.Product
	id := c.Param("id")

	if err := db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.Name = input.Name
	product.CategoryID = input.CategoryID
	product.Stok = input.Stok
	product.Harga = input.Harga
	product.UpdatedAt = time.Now()

	db.Save(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// DELETE
func DeleteProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var product models.Product
	id := c.Param("id")

	if err := db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	db.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
