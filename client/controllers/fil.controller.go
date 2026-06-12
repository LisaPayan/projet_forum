package controllers

import (
	"client/services"
	"client/templates"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type FilControllers struct {
	service  *services.FilService
	template *templates.TemplateManager
}

// InitProductController crée un contrôleur produit avec son service et son moteur de templates.
func InitFilController(service *services.FilService, template *templates.TemplateManager) *FilControllers {
	return &FilControllers{service: service, template: template}
}

func (c *FilControllers) DisplayList(w http.ResponseWriter, r *http.Request) {
	filsList, filsErr := c.service.ReadAll()
	if filsErr != nil {
		http.Error(w, filsErr.Error(), http.StatusInternalServerError)
		return
	}
	c.template.RenderTemplate(w, r, "menu", filsList)
	fmt.Println(filsList)
}

func (c *FilControllers) DisplayMessagesFil(w http.ResponseWriter, r *http.Request) {
	idFil, idFilErr := strconv.Atoi(mux.Vars(r)["id"])
	if idFilErr != nil {
		http.Error(w, "Erreur - Identifiant produit invalide", http.StatusBadRequest)
		return
	}

	fil, filErr := c.service.ReadByIdMessages(idFil)
	if filErr != nil {
		http.Error(w, filErr.Error(), http.StatusInternalServerError)
		return
	}

	c.template.RenderTemplate(w, r, "details_fil", fil)
}

func (c *FilControllers) DisplayListPetanque(w http.ResponseWriter, r *http.Request) {
	filList, filErr := c.service.FilsPetanque()
	if filErr != nil {
		http.Error(w, filErr.Error(), http.StatusInternalServerError)
		return
	}
	c.template.RenderTemplate(w, r, "petanque", filList)
}

func (c *FilControllers) DisplayListCuisine(w http.ResponseWriter, r *http.Request) {
	filList, filErr := c.service.FilsCuisine()
	if filErr != nil {
		http.Error(w, filErr.Error(), http.StatusInternalServerError)
		return
	}
	c.template.RenderTemplate(w, r, "cuisine", filList)
}

func (c *FilControllers) DisplayListNature(w http.ResponseWriter, r *http.Request) {
	filList, filErr := c.service.FilsNature()
	if filErr != nil {
		http.Error(w, filErr.Error(), http.StatusInternalServerError)
		return
	}
	c.template.RenderTemplate(w, r, "nature", filList)
}
