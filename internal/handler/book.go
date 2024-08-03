package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
	// TODO: Implement get all books logic
	c.JSON(http.StatusOK, gin.H{"message": "Get all books"})
}

func CreateBook(c *gin.Context) {
	var input struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
		ImageURL    string `json:"image_url"`
		ReleaseYear int    `json:"release_year" binding:"required,min=1980,max=2024"`
		Price       int    `json:"price" binding:"required"`
		TotalPage   int    `json:"total_page" binding:"required"`
		CategoryID  int    `json:"category_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Calculate thickness
	thickness := "tipis"
	if input.TotalPage > 100 {
		thickness = "tebal"
	}

	// TODO: Save book to database
	c.JSON(http.StatusCreated, gin.H{"message": "Book created", "thickness": thickness})
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	// TODO: Implement get book logic
	c.JSON(http.StatusOK, gin.H{"message": "Get book " + id})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	// TODO: Implement delete book logic
	// Check if book exists before deleting
	c.JSON(http.StatusOK, gin.H{"message": "Book " + id + " deleted"})
}
