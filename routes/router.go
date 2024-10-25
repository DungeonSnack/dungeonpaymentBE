package routes

import (
	"dungeonSnackBE/controller/profil"
	"dungeonSnackBE/controller/auth"
	"dungeonSnackBE/controller/produk"
	"github.com/gorilla/mux"
)

// InitializeRoutes sets up the router
func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()
	// Ke kontroller auth
	router.HandleFunc("/registrasi", auth.Register).Methods("POST")
	router.HandleFunc("/login", auth.Login).Methods("POST")
	//ke kontroller profil
	router.HandleFunc("/profil-update", profil.UpdateProfil).Methods("PUT")
	router.HandleFunc("/profil", profil.GetProfil).Methods("GET")
	router.HandleFunc("/produk", produk.GetProduk).Methods("GET")
	router.HandleFunc("/produk", produk.AddProduk).Methods("POST")
	router.HandleFunc("/produk/{id}", produk.UpdateProduk).Methods("PUT")
	router.HandleFunc("/produk/{id}", produk.DeleteProduk).Methods("DELETE")
	return router
}