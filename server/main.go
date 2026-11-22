package main

import (
	"log"
	"net/http"
	"projet_forum/app"
	"projet_forum/config"
	"projet_forum/middleware"
)

func main() {
	app := app.InitApp()
	defer app.Close()

	port := config.GetEnvWithDefault("PORT", "8080")

	// Lancement du serveur
	log.Printf("Serveur lancé : http://localhost:%s", port)
	serveErr := http.ListenAndServe(":"+port, middleware.LoggingMiddleware(app.Router))
	if serveErr != nil {
		log.Fatalf("Erreur lancement serveur - %s", serveErr.Error())
	}
}
