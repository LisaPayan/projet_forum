package routers

import (
	"client/controllers"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(r *mux.Router, filController *controllers.FilControllers) {
	r.HandleFunc("/", filController.DisplayList).Methods("GET")
	r.HandleFunc("/fils/petanque", filController.DisplayListPetanque).Methods("GET")
	r.HandleFunc("/fils/cuisine", filController.DisplayListCuisine).Methods("GET")
	r.HandleFunc("/fils/nature", filController.DisplayListNature).Methods("GET")
	r.HandleFunc("/fil/{id}/messages", filController.DisplayMessagesFil).Methods("GET")

	// r.HandleFunc("/product/create", filController.CreateForm).Methods("GET")
	// r.HandleFunc("/product", filController.Create).Methods("POST")
}
