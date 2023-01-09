package main

import (
	"log"
	"webapi-gingorm/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/gingorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error")
	}

	// db.AutoMigrate(&book.Book{})

	//CREATE
	//STRUCT
	// book := book.Book{}
	// book.Title = "Naruto Shippuden"
	// book.Price = 120000
	// book.Discount = 20
	// book.Rating = 4
	// book.Description = "ini adalah buku komik Naruto Shippuden"

	// err = db.Create(&book).Error //untuk menyimpan
	// if err != nil {
	// 	fmt.Println("Error Creating Book")
	// }

	//Menampil suatu data /satu data
	// var book book.Book
	// err = db.Debug().First(&book, 6).Error // [First] menampilkan data pertama (awal). [Last] menampilkan data terakhir
	// if err != nil {
	// 	fmt.Println("Error Finding Book")
	// }
	// fmt.Println("Title :", book.Title)
	// fmt.Println("book object %v", book)

	//menampilkan banyak data
	// var books []book.Book
	// err = db.Debug().Find(&books).Error
	// if err != nil {
	// 	fmt.Println("Error Finding Book")
	// }
	// for _, b := range books {
	// 	fmt.Println("Title :", b.Title)
	// 	fmt.Println("book object %v", b)
	// }

	//mencari buku
	// var books []book.Book
	// err = db.Debug().Where("rating =?", 5).Find(&books).Error
	// if err != nil {
	// 	fmt.Println("Error Finding Book")
	// }
	// for _, b := range books {
	// 	fmt.Println("Title :", b.Title)
	// 	fmt.Println("book object %v", b)
	// }

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RouteHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run(":8000")

}
