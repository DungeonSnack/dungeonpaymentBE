package routes

import (
	"dungeonSnackBE/controller/auth"
	"dungeonSnackBE/controller/order"
	"dungeonSnackBE/controller/pembayaran"
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

	router.HandleFunc("/order/{slug}/pembayaran", pembayaran.AddpembayaranToorder).Methods("POST")
	router.HandleFunc("/order/{slug}/pembayaran", pembayaran.GetpembayaranByorder).Methods("GET")
	router.HandleFunc("/pembayaran-id", pembayaran.GetpembayaranByID).Methods("GET")
	router.HandleFunc("/order/pembayaran/update", pembayaran.Updatepembayaran).Methods("PUT")
	router.HandleFunc("/order/{slug}/pembayaran", pembayaran.Deletepembayaran).Methods("DELETE")

	return router
}
