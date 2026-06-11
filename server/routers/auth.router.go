package routers

import (
	"projet_forum/controllers"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router, authController *controllers.AuthController) {
	r.HandleFunc("/register", authController.Register).Methods("POST")
}
