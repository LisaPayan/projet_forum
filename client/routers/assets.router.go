package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterAssetsRoutes(r *mux.Router) {
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
}
