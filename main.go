package main

import (
	"dungeonSnackBE/config"
	"dungeonSnackBE/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	// Create a new router instance
	router := mux.NewRouter()

	// Rute untuk login
	router.HandleFunc("/registrasi", func(w http.ResponseWriter, r *http.Request) {
		controller.Register(w, r, connectdb.Client())
	}).Methods("POST")
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		controller.Login(w, r, connectdb.Client())
	}).Methods("POST")

	// Menjalankan server pada port 8080
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
