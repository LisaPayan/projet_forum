// Package routers configure les routes de l'application.
package routers

import (
	"client/controllers"

	"github.com/gorilla/mux"
)

// RegisterProductRoutes enregistre les routes liées aux produits pour l'interface web.
func RegisterCookiesRoutes(r *mux.Router, cookiesController *controllers.CookiesControllers) {
	r.HandleFunc("/cookies/add", controllers.AddCoockieHandler).Methods("GET")
	r.HandleFunc("/cookies/read", controllers.ReadCookieHandler).Methods("GET")
	r.HandleFunc("/cookies/delete", controllers.DeleteCookieHandler).Methods("GET")
}
