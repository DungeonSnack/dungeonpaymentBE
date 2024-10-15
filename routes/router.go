package routes

import (
	"dungeonSnackBE/controller/auth"
	"dungeonSnackBE/controller/kategori"


	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

// InitializeRoutes sets up the router
func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()
	
	// Define your routes here
	router.HandleFunc("/regis", auth.Register).Methods("POST")
	router.HandleFunc("/login", auth.Login).Methods("POST")


	return router
}

func NewKategoriRouter(db *mongo.Client) *mux.Router {
	router := mux.NewRouter()

	// Route untuk membuat kategori baru
	router.HandleFunc("/kategori", func(w http.ResponseWriter, r *http.Request) {
		controller.CreateKategori(w, r, db)
	}).Methods("POST")

	// Route untuk mendapatkan semua kategori
	router.HandleFunc("/kategori", func(w http.ResponseWriter, r *http.Request) {
		controller.GetKategori(w, r, db)
	}).Methods("GET")

	// Route untuk menghapus kategori berdasarkan ID
	router.HandleFunc("/kategori", func(w http.ResponseWriter, r *http.Request) {
		controller.DeleteKategori(w, r, db)
	}).Methods("DELETE")

	// Route untuk mengupdate kategori
	router.HandleFunc("/kategori", func(w http.ResponseWriter, r *http.Request) {
		controller.UpdateKategori(w, r, db)
	}).Methods("PUT")

	return router
}