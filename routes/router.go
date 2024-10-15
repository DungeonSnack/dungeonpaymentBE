package routes

import (
	"dungeonSnackBE/controller/auth"
	"dungeonSnackBE/controller/kategori" // Ensure this path is correct and the package exists
	"dungeonSnackBE/controller/profil"
	"github.com/gorilla/mux"
)

// InitializeRoutes sets up the router
func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()
	
	// Define your routes here
	router.HandleFunc("/regis", auth.Register).Methods("POST")
	router.HandleFunc("/login", auth.Login).Methods("POST")

	router.HandleFunc("/kategori", kategori.CreateKategori).Methods("POST")
	router.HandleFunc("/kategori", kategori.GetKategori).Methods("GET")
	router.HandleFunc("/kategori/{id}", kategori.UpdateKategori).Methods("PUT")
	router.HandleFunc("/kategori/{id}", kategori.DeleteKategori).Methods("DELETE")

	router.HandleFunc("/profil", profil.GetProfile).Methods("GET")
	router.HandleFunc("/profil", profil.UpdateProfile).Methods("PUT")


	return router
}