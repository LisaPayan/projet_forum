package routers

import (
	"client/controllers"
	"client/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(r *mux.Router, filController *controllers.FilControllers) {
	r.HandleFunc("/", filController.DisplayList).Methods("GET")
	r.HandleFunc("/all", filController.DisplayPaginationAll).Methods("GET")
	r.Handle("/search", middleware.AuthMiddleware(http.HandlerFunc(filController.DisplaySearch))).Methods("GET")

	r.HandleFunc("/fils/petanque", filController.DisplayPaginationPetanque).Methods("GET")
	r.HandleFunc("/fils/cuisine", filController.DisplayPaginationCuisine).Methods("GET")
	r.HandleFunc("/fils/nature", filController.DisplayPaginationNature).Methods("GET")

	r.HandleFunc("/fil/{id}/messages", filController.DisplayPaginationMessage).Methods("GET")
	r.Handle("/fil/{id}/messages", middleware.AuthMiddleware(http.HandlerFunc(filController.CreateMessage))).Methods("POST")

	r.Handle("/fils/create", middleware.AuthMiddleware(http.HandlerFunc(filController.CreateForm))).Methods("GET")
	r.Handle("/fils", middleware.AuthMiddleware(http.HandlerFunc(filController.Create))).Methods("POST")

}
