package routers

import (
	"github.com/gorilla/mux"
)

func SetupUserRoutes(r *mux.Router) {
	//users := r.PathPrefix("/users").Subrouter()

	//users.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	//users.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	// Protected routes
	//users.Use(middleware.Auth)
	//users.HandleFunc("/profile", handlers.ProfileHandler).Methods("GET")
	//users.HandleFunc("/profile", handlers.UpdateProfileHandler).Methods("PUT")
}
