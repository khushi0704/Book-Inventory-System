package main

import (
	"fmt"
	"gofr.dev/pkg/gofr"
	"book-inventory-system/controllers"
)


func main(){
	
	// goFr instance 
	app := gofr.New()

	// createBook - > http://localhost:8000/books
	// /books-> a book is getting created 
	app.POST("/books",controllers.CreateBook)

	// getBook 
	app.GET("/book",controllers.GetBook)

	// getAll
	app.GET("/books",controllers.GetAllBooks)

	// updateBook -
	// partial update 
	app.PATCH("/book/{id}",controllers.UpdateBook)

    // Starts the server, it will listen on the default port 8000.
    // it can be over-ridden through configs
	fmt.Println("Starting server ..... ")
    app.Start()
	fmt.Println("Server started.")
}