package routes

import (
	"dungeonSnackBE/controller/auth"
	"dungeonSnackBE/controller/kategori" // Ensure this path is correct and the package exists
	"dungeonSnackBE/controller/produk"
	"dungeonSnackBE/controller/profil"
	"dungeonSnackBE/controller/toko"

	"github.com/gorilla/mux"
)

// InitializeRoutes sets up the router
func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()
	// Define your routes here
	router.HandleFunc("/regis", auth.Register).Methods("POST")
	router.HandleFunc("/login", auth.Login).Methods("POST")

	router.HandleFunc("/profil", profil.GetProfile).Methods("GET")
	router.HandleFunc("/profil/update", profil.UpdateProfile).Methods("PUT")

	router.HandleFunc("/addKategori", kategori.AddKategori).Methods("POST")
	router.HandleFunc("/kategori", kategori.GetKategori).Methods("GET")
	router.HandleFunc("/kategori", kategori.DeleteKategori).Methods("DELETE")
	router.HandleFunc("/kategori", kategori.Updatekategori).Methods("PUT")

	router.HandleFunc("/kategori/produk", produk.AddProduk).Methods("POST")
	router.HandleFunc("/kategori/{slug}/produk", produk.GetProdukByKategori).Methods("GET")
	router.HandleFunc("/kategori/produk/update", produk.UpdateProduk).Methods("PUT")
	router.HandleFunc("/kategori/{slug}/produk", produk.DeleteProduk).Methods("DELETE")

	router.HandleFunc("/toko", toko.GetToko).Methods("GET")
	router.HandleFunc("/toko", toko.AddToko).Methods("POST")
	router.HandleFunc("/toko/update", toko.UpdateToko).Methods("PUT")
	router.HandleFunc("/toko/delete", toko.DeleteToko).Methods("DELETE")
	return router
}
