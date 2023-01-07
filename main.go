package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", routeHandler)
	router.GET("/hello", helloHandler)
	router.GET("/books/:id", booksHandler)
	router.GET("/query", queryHandler)
	router.POST("/query", queryHandler)

	router.Run(":8000")

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

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{"id": title, "price": price})
}
