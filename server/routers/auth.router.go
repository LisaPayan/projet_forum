package routers

import (
	"net/http"
	"projet_forum/controllers"
	"projet_forum/middleware"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router, authController *controllers.AuthController) {
	r.HandleFunc("/register", authController.Register).Methods("POST")

	r.HandleFunc("/login", authController.Login).Methods("POST")

	r.Handle("/me", middleware.AuthMiddleware(http.HandlerFunc(authController.Me))).Methods("GET")
}
