package controllers

import (
	"client/services"
	"client/templates"
	"net/http"
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
	productList, productErr := c.service.ReadAll()
	if productErr != nil {
		http.Error(w, productErr.Error(), http.StatusInternalServerError)
		return
	}
	c.template.RenderTemplate(w, r, "menu", productList)
}
