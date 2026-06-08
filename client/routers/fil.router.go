package routers

import (
	"client/controllers"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(r *mux.Router, filController *controllers.FilControllers) {
	r.HandleFunc("/", filController.DisplayList).Methods("GET")
	// r.HandleFunc("/product/create", filController.CreateForm).Methods("GET")
	// r.HandleFunc("/product/{id}", filController.DisplayById).Methods("GET")
	// r.HandleFunc("/product", filController.Create).Methods("POST")
}
