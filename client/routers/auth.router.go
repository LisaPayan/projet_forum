package routers

import (
	"client/controllers"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(r *mux.Router, authController *controllers.AuthControllers) {
	// Route GET : affiche le formulaire de connexion.
	r.HandleFunc("/login", authController.LoginForm).Methods("GET")

	// Route POST : traite les donnees envoyees par le formulaire de connexion.
	r.HandleFunc("/login", authController.Login).Methods("POST")

	r.HandleFunc("/logout", authController.Logout).Methods("GET")
	r.HandleFunc("/me", authController.Me).Methods("GET")

	r.HandleFunc("/register", authController.RegisterForm).Methods("GET")
	r.HandleFunc("/register", authController.Register).Methods("POST")

}
