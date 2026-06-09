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

	// Initilisation des repositories
	filRepository := repositories.InitFilRepository(db)

	// Initilisation des services
	filService := services.InitFilService(filRepository)

	// Initilisation des controllers
	filController := controllers.InitFilController(filService)

	// Enregistrement des routes (avec ajout du préfix "/api/...")
	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	routers.RegisterFilRoutes(router, filController)

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
