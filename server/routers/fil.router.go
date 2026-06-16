package routers

import (
	"net/http"
	"projet_forum/controllers"
	"projet_forum/middleware"

	"github.com/gorilla/mux"
)

func RegisterFilRoutes(r *mux.Router, filController *controllers.FilControllers) {
	r.HandleFunc("/fils", filController.ReadAll).Methods("GET")
	r.HandleFunc("/fil/{id}", filController.ReadById).Methods("GET")
	r.Handle("/fils", middleware.AuthMiddleware(http.HandlerFunc(filController.CreateFil))).Methods("POST")
	r.HandleFunc("/fils/petanque", filController.FilsPetanque).Methods("GET")
	r.HandleFunc("/fils/cuisine", filController.FilsCuisine).Methods("GET")
	r.HandleFunc("/fils/nature", filController.FilsNature).Methods("GET")
	r.HandleFunc("/fil/{id}/messages", filController.GetMessagesByFil).Methods("GET")
	r.Handle("/fil/{id}/messages", middleware.AuthMiddleware(http.HandlerFunc(filController.CreateMessage))).Methods("POST")

	// r.HandleFunc("/fil/{id}/messages?ord=asc", filController.GetMessagesByFilAnciens).Methods("GET")
	// r.HandleFunc("/fil/{id}/messages?ord=desc", filController.GetMessagesByFilRecents).Methods("GET")

}
