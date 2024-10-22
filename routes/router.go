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

	router.HandleFunc("/profil/update", profil.UpdateProfile).Methods("POST")
	router.HandleFunc("/profil", profil.GetProfile).Methods("GET")
	return router
}
