package routers

import (
	"projet_forum/controllers"

	"github.com/gorilla/mux"
)

func RegisterFilRoutes(r *mux.Router, filController *controllers.FilControllers) {
	r.HandleFunc("/fils", filController.ReadAll).Methods("GET")
	r.HandleFunc("/fil/{id}/messages", filController.GetMessagesByFil).Methods("GET")

}
