package controllers

import (
	"client/services"
	"client/templates"
)

type AuthControllers struct {
	service  *services.AuthService
	template *templates.TemplateManager
}

func InitAuthController(service *services.AuthService, template *templates.TemplateManager) *AuthControllers {
	return &AuthControllers{service: service, template: template}
}
