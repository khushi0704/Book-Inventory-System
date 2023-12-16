package main

import (
	"fmt"
	"gofr.dev/pkg/gofr"
	"book-inventory-system/controllers"
)

func main(){
	
	// goFr instance 
	app := gofr.New()

	// /books-> a book is getting created 
	app.POST("/books",controllers.CreateBook)

	// to get a book info 
	app.GET("/book",controllers.GetBook)

	// to retreive list of all books
	app.GET("/books",controllers.GetAllBooks)

	// partial update 
	// to update a book by its quantity available in stock
	app.PATCH("/book",controllers.UpdateBookQuantity)

	// to remove a book from db 
	app.DELETE("/book",controllers.DeleteBook)

    // Starts the server, it will listen on the default port 8000.
    // it can be over-ridden through configs
	fmt.Println("Starting server ..... ")
    app.Start()
	fmt.Println("Server started.")
}