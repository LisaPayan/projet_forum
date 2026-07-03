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

	r.Handle("/message/reaction", middleware.AuthMiddleware(http.HandlerFunc(filController.AjoutReaction))).Methods("POST")

	r.Handle("/fils/create", middleware.AuthMiddleware(http.HandlerFunc(filController.CreateForm))).Methods("GET")
	r.Handle("/fils", middleware.AuthMiddleware(http.HandlerFunc(filController.Create))).Methods("POST")

	r.Handle("/fil/{id}/edit", middleware.AuthMiddleware(http.HandlerFunc(filController.UpdateForm))).Methods("GET")
	r.Handle("/fil/{id}/update", middleware.AuthMiddleware(http.HandlerFunc(filController.UpdateFilById))).Methods("POST")
	r.Handle("/fil/{id}/delete", middleware.AuthMiddleware(http.HandlerFunc(filController.DeleteFilById))).Methods("POST")

	r.Handle("/message/edit", middleware.AuthMiddleware(http.HandlerFunc(filController.UpdateFormMessage))).Methods("GET")
	r.Handle("/message/update", middleware.AuthMiddleware(http.HandlerFunc(filController.UpdateMessageById))).Methods("POST")
	r.Handle("/message/delete", middleware.AuthMiddleware(http.HandlerFunc(filController.DeleteMessage))).Methods("POST")

	r.Handle("/dashboard", middleware.AuthMiddleware(http.HandlerFunc(filController.DashboardHandler))).Methods("GET")
}
