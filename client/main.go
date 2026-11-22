package main

import (
	"client/app"
	"client/config"
	"client/middleware"
	"log"
	"net/http"
)

func main() {
	app := app.InitApp()
	port := config.GetEnvWithDefault("PORT", "3000")

	log.Printf("Serveur lancé : http://localhost:%s", port)
	serveErr := http.ListenAndServe(":"+port, middleware.LoggingMiddleware(app.Router))
	if serveErr != nil {
		log.Fatalf("Erreur lancement serveur - %s", serveErr.Error())
	}
}
