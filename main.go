package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func main() {
	router := gin.Default()

	router.GET("/", routeHandler)
	router.GET("/hello", helloHandler)
	router.GET("/books/:id", booksHandler)
	router.GET("/query", queryHandler)
	router.POST("/books", postBooksHandler)

	router.Run(":8000")

}

type BookInput struct {
	Title string `json:"title" binding:"required"`
	Price int    `json:"price" binding:"required,number"`
}

func routeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Ringin Restu Pati",
		"bio":  "Software Engineer",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Restu Pati",
		"bio":  "Information System",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{"id": id})
}
func postBooksHandler(c *gin.Context) {
	var bookInput BookInput

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

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{"id": title, "price": price})
}
