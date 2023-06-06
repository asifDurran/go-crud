package main

import (
	"crud-golang/config"
	"crud-golang/controller"
	"crud-golang/helper"
	"crud-golang/repository"
	"crud-golang/router"
	"crud-golang/service"
	"fmt"
	"net/http"
)


func main()  {
	fmt.Printf("Server start")
	// Database
	db := config.DatabaseConnection()

	// repository
	bookRepository := repository.NewBookRepository(db)

	// service
	bookService := service.NewBookServiceImpl(bookRepository)

	// controller
	bookController := controller.NewBookController(bookService)

    
	// routes 
	routes := router.NewRouter(bookController)

	server := http.Server{Addr: "localhost:8888", Handler: routes}

	err := server.ListenAndServe()

	helper.PanicIfError(err)
}