package routers

import (
	"client/controllers"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(r *mux.Router, filController *controllers.FilControllers) {
	r.HandleFunc("/", filController.DisplayList).Methods("GET")
	r.HandleFunc("/all", filController.DisplayPaginationAll).Methods("GET")
	r.HandleFunc("/search", filController.DisplaySearch).Methods("GET")
	r.HandleFunc("/fils/petanque", filController.DisplayPaginationPetanque).Methods("GET")
	r.HandleFunc("/fils/cuisine", filController.DisplayPaginationCuisine).Methods("GET")
	r.HandleFunc("/fils/nature", filController.DisplayPaginationNature).Methods("GET")
	r.HandleFunc("/fil/{id}/messages", filController.DisplayPaginationMessage).Methods("GET")

	// r.HandleFunc("/product/create", filController.CreateForm).Methods("GET")
	// r.HandleFunc("/product", filController.Create).Methods("POST")
}
