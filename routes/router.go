package routes

import (
	"dungeonSnackBE/controller/auth"
	"dungeonSnackBE/controller/order"
	"dungeonSnackBE/controller/pembayaran"
	"dungeonSnackBE/controller/produk"
	"dungeonSnackBE/controller/profil"

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

	router.HandleFunc("/order", order.Getorder).Methods("GET")
	router.HandleFunc("/add-order", order.Addorder).Methods("POST")
	router.HandleFunc("/order-id", order.GetorderByID).Methods("GET")
	router.HandleFunc("/order/update", order.UpdateorderByID).Methods("PUT")
	router.HandleFunc("/order/delete", order.DeleteorderByID).Methods("DELETE")

	router.HandleFunc("/order/pembayaran", pembayaran.AddpembayaranToorder).Methods("POST")
	router.HandleFunc("/order/{slug}/pembayaran", pembayaran.GetpembayaranByorder).Methods("GET")
	router.HandleFunc("/pembayaran-id", pembayaran.GetpembayaranByID).Methods("GET")
	router.HandleFunc("/order/pembayaran/update", pembayaran.Updatepembayaran).Methods("PUT")
	router.HandleFunc("/order/{slug}/pembayaran", pembayaran.Deletepembayaran).Methods("DELETE")

	router.HandleFunc("/produk", produk.CreateProduk).Methods("POST")
	router.HandleFunc("/produk", produk.GetProduk).Methods("GET")
	router.HandleFunc("/produk/{slug}", produk.GetProdukBySlug).Methods("GET")
	router.HandleFunc("/produk/{slug}", produk.GetProdukByID).Methods("GET")
	router.HandleFunc("/produk/{slug}", produk.GetProdukByCategory).Methods("GET")
	router.HandleFunc("/produk/update", produk.UpdateProduk).Methods("PUT")
	router.HandleFunc("/produk/delete", produk.DeleteProduk).Methods("DELETE")

	return router
}
