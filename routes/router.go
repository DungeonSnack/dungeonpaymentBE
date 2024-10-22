package routes

import (
	"dungeonSnackBE/controller/profil"
	"dungeonSnackBE/controller/auth"

	"github.com/gorilla/mux"
)

// InitializeRoutes sets up the router
func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()
	// Define your routes here
	router.HandleFunc("/regis", auth.Register).Methods("POST")
	router.HandleFunc("/login", auth.Login).Methods("POST")

	router.HandleFunc("/profil/update", profil.UpdateProfil).Methods("PUT")
	router.HandleFunc("/profil", profil.GetProfil).Methods("GET")
	return router
}