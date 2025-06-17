package main

import (
	"booknest-api/database"
	"booknest-api/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Initialize database connection
	db := database.Connect()
	defer db.Close()

	// Setup routes
	r := routes.SetupRoutes(db)

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
} 