package routes

import (
	"dungeonSnackBE/controller/auth"
	"dungeonSnackBE/controller/kategori" // Ensure this path is correct and the package exists
	"dungeonSnackBE/controller/profil"
	"dungeonSnackBE/controller/produk"
	"github.com/gorilla/mux"
)

// InitializeRoutes sets up the router
func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()
	// Define your routes here
	router.HandleFunc("/regis", auth.Register).Methods("POST")
	router.HandleFunc("/login", auth.Login).Methods("POST")

	router.HandleFunc("/kategori", kategori.AddKategori).Methods("POST")

	router.HandleFunc("/profil", profil.GetProfile).Methods("GET")
	router.HandleFunc("/profil", profil.UpdateProfile).Methods("PUT")

	router.HandleFunc("/produk", produk.AddProduk).Methods("POST")


	return router
}