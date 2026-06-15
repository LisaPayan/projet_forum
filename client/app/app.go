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
	authApi := api.InitAuthApi(baseURL)

	filService := services.InitFilService(filApi)
	authService := services.InitAuthService(authApi)

	filController := controllers.InitFilController(filService, templatesManager)
	authController := controllers.InitAuthController(authService, templatesManager)

	router := mux.NewRouter()
	routers.RegisterAssetsRoutes(router)
	routers.RegisterProductRoutes(router, filController)
	routers.AuthClientRoutes(router, authController)

	return &App{
		Router: router,
	}
}
