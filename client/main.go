package main

import (
	"client/app"
	"log"
	"net/http"
)

func main() {
	app := app.InitApp()

	log.Printf("Serveur lancé : http://localhost:3000")
	serveErr := http.ListenAndServe(":3000", app.Router)
	if serveErr != nil {
		log.Fatalf("Erreur lancement serveur - %s", serveErr.Error())
	}
}
