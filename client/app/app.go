package app

import (
	"client/api"
	"client/config"
	"client/controllers"
	"client/routers"
	"client/services"
	"client/templates"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

// InitApp initialise l'application, charge les templates et configure le routeur.
func InitApp() *App {
	config.LoadEnv()

	templatesManager := templates.NewTemplatesManager()

	baseURL := config.GetRequiredEnv("BASE_URL")
	filApi := api.InitFilApi(baseURL)

	filService := services.InitFilService(filApi)

	filController := controllers.InitFilController(filService, templatesManager)

	router := mux.NewRouter()
	routers.RegisterAssetsRoutes(router)
	routers.RegisterProductRoutes(router, filController)

	return &App{
		Router: router,
	}
}
