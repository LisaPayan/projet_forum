package app

import (
	"database/sql"
	"projet_forum/config"
	"projet_forum/controllers"
	"projet_forum/repositories"
	"projet_forum/routers"
	"projet_forum/services"

	"github.com/gorilla/mux"
)

type App struct {
	Db     *sql.DB
	Router *mux.Router
}

func InitApp() *App {
	config.LoadEnv()

	db := config.InitDB()

	// Initialisation des repositories
	filRepository := repositories.InitFilRepository(db)
	authRepository := repositories.InitAuthRepository(db)

	// Initialisation des services
	filService := services.InitFilService(filRepository)
	authService := services.InitAuthService(authRepository)

	// Initialisation des controllers
	filController := controllers.InitFilController(filService)
	authController := controllers.InitAuthController(authService)

	// Enregistrement des routes avec le prefixe /api
	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	routers.RegisterFilRoutes(router, filController)
	routers.AuthRoutes(router, authController)

	return &App{
		Db:     db,
		Router: router,
	}
}

func (a *App) Close() {
	if a.Db != nil {
		a.Db.Close()
	}
}
