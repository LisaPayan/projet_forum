package routers

import (
	"projet_forum/controllers"

	"github.com/gorilla/mux"
)

func RegisterFilRoutes(r *mux.Router, filController *controllers.FilControllers) {
	r.HandleFunc("/fils", filController.ReadAll).Methods("GET")
	r.HandleFunc("/fils/petanque", filController.FilsPetanque).Methods("GET")
	r.HandleFunc("/fils/cuisine", filController.FilsCuisine).Methods("GET")
	r.HandleFunc("/fils/nature", filController.FilsNature).Methods("GET")
	r.HandleFunc("/fil/{id}/messages", filController.GetMessagesByFil).Methods("GET")
	r.HandleFunc("/fil/{id}/messages?ord=asc", filController.GetMessagesByFilAnciens).Methods("GET")
	r.HandleFunc("/fil/{id}/messages?ord=desc", filController.GetMessagesByFilRecents).Methods("GET")

}
