package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"neohub.asia/mod/databases/models"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// GET /book
func GetBooks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var books []models.Book
	db.Find(&books)
	c.Set("response", books)
}

// GET /book/:id
func GetBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.Set("error", map[string]interface{}{
			"code": http.StatusBadRequest,
			"data": "Book not found!",
		})
		return
	}
	c.Set("response", book)
}

// POST /book
func CreateBook(c *gin.Context) {
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Set("error", map[string]interface{}{
			"code": http.StatusUnprocessableEntity,
			"data": err.Error(),
		})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Create(&book).Error; err != nil {
		c.Set("error", map[string]interface{}{
			"code": http.StatusBadRequest,
			"data": "Could not create book!",
		})
		return
	}

	c.Set("response", book)
	c.Set("status", http.StatusCreated)
}

// PATCH /book/:id
func UpdateBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var book models.Book

	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.Set("error", map[string]interface{}{
			"code": http.StatusBadRequest,
			"data": "Book not found!",
		})
		return
	}

	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Set("error", map[string]interface{}{
			"code": http.StatusUnprocessableEntity,
			"data": err.Error(),
		})
		return
	}

	updateData := map[string]interface{}{}
	if input.Title != "" {
		updateData["title"] = input.Title
	}
	if input.Author != "" {
		updateData["author"] = input.Author
	}

	if err := db.Model(&book).Updates(updateData).Error; err != nil {
		c.Set("error", map[string]interface{}{
			"code": http.StatusBadRequest,
			"data": "Failed to update book!",
		})
		return
	}

	c.Set("response", book)
}

// DELETE /book/:id
func DeleteBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var book models.Book

	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.Set("error", map[string]interface{}{
			"code": http.StatusBadRequest,
			"data": "Book not found!",
		})
		return
	}

	if err := db.Delete(&book).Error; err != nil {
		c.Set("error", map[string]interface{}{
			"code": http.StatusBadRequest,
			"data": "Failed to delete book!",
		})
		return
	}

	c.Set("response", book)
}
