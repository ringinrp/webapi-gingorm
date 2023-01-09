package handler

import (
	"fmt"
	"net/http"

	"webapi-gingorm/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func RouteHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Ringin Restu Pati",
		"bio":  "Software Engineer",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Restu Pati",
		"bio":  "Information System",
	})
}

func BooksHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{"id": id})
}
func PostBooksHandler(c *gin.Context) {
	var bookInput book.BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Title": bookInput.Title,
		"Price": bookInput.Price,
	})
}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{"id": title, "price": price})
}
