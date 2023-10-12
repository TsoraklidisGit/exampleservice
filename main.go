package main

import (
	"exampleservice/controller"
	"exampleservice/repository"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Initialize the database
	db, err := repository.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the controller with the database connection
	controllers := controller.NewController(db)

	// Register routes and handlers
	http.HandleFunc("/createItem", controllers.CreateItem)
	http.HandleFunc("/getItems", controllers.GetItems)
	http.HandleFunc("/items/{id}", controllers.UpdateItem)
	http.HandleFunc("/deleteItem/{id}", controllers.DeleteItem)

	port := ":8080"
	fmt.Printf("Server started on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
