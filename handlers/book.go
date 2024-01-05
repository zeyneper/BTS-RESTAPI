package handlers

import (
	"mentalartsapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetBooks      godoc
// @Summary      Get all books array
// @Description  Responds with the list of all books as JSON
// @Tags         Books
// @Produce      json
// @Success      200  {array}  models.Book
// @Failure      400  {object} nil
// @Router       /book [get]
func (h Handler) GetBooks(c *gin.Context) {
	var books []models.Book

	if err := h.db.Preload("User").Find(&books).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, books)
}

// GetBookById      godoc
// @Summary      Get book by Id
// @Description  Returns the book whit specific ID
// @Tags         Books
// @Produce      json
// @Success      200  {object}  models.Book
// @Failure      400  {object} nil
// @Router       /book/{id} [get]
func (h Handler) GetBookById(c *gin.Context) {
	id := c.Params.ByName("id")
	var book models.Book
	if err := h.db.Where("id = ?", id).First(&book).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

// CreateBook      godoc
// @Summary      Create a new book JSON and stores in DB. Returns saved JSON.
// @Description  Returns the book whit specific ID
// @Tags         Books
// @Param        book body models.Book true "Book JSON"
// @Produce      json
// @Success      201  {object}  models.Book
// @Failure      400  {object} nil
// @Router       /book [post]
func (h Handler) UpdateBook(c *gin.Context) {
	var book models.Book

	err := c.BindJSON(&book)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}
	book.ID = 0

	if err := h.db.Create(&book).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

// DeleteBookById      godoc
// @Summary      Delete book by Id
// @Description  Deletes book with specific Id
// @Tags         Books
// @Param        id path int true "delete book by Id"
// @Produce      json
// @Success      200
// @Failure      400  {object} nil
// @Router       /book/{id} [delete]
func (h Handler) DeleteBookById(c *gin.Context) {
	id := c.Params.ByName("id")

	if err := h.db.Where("id = ?", id).Delete(&models.Book{}).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
	}
}
