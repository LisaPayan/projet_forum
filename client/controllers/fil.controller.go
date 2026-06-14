package controllers

import (
	"client/dto"
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

func (c *FilControllers) DisplayPagination(w http.ResponseWriter, r *http.Request) {
	pageStr := r.FormValue("page")
	pageInt, _ := strconv.Atoi(pageStr)

	if pageInt < 0 {
		pageInt = 0
	}

	startIndex := pageInt * 10
	endIndex := startIndex + 10

	data, err := c.service.ReadAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if startIndex >= len(data) {
		pageInt = 0
		startIndex = 0
		endIndex = 10
	}

	if endIndex > len(data) {
		endIndex = len(data)
	}

	SelectCountries := data[startIndex:endIndex]

	prevPage := pageInt - 1
	if prevPage < 0 {
		prevPage = 0
	}

	nextPage := pageInt
	if endIndex < len(data) {
		nextPage = pageInt + 1
	}

	vieData := dto.PagePagination{
		Page: pageInt,
		Next: nextPage,
		Prev: prevPage,
		Data: SelectCountries,
	}

	c.template.RenderTemplate(w, r, "all", vieData)

}
