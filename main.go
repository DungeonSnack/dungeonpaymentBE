package main

import (
	"dungeonSnackBE/config"
	"dungeonSnackBE/routes"
	"fmt"
	"log"
	"net/http"
	"github.com/rs/cors"
)

func main() {
	// Hubungkan ke database MongoDB
	connectdb := config.Mongoconn

	if config.ErrorMongoconn != nil {
		fmt.Println("Failed to connect to MongoDB:", config.ErrorMongoconn)
		return
	}

	// Check if the connection is successful
	if connectdb != nil {
		fmt.Println("Successfully connected to MongoDB!")
	} else {
		fmt.Println("MongoDB connection is nil")
	}

	router := routes.InitializeRoutes()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            true,
	})

	handler := c.Handler(router)
	// Menjalankan server pada port 8080
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
