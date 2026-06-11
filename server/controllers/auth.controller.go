package controllers

import (
	"encoding/json"
	"net/http"
	"projet_forum/dto"
	"projet_forum/helper"
	"projet_forum/services"
)

type AuthController struct {
	service *services.AuthService
}

func InitAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{service: authService}
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var data dto.RegisterRequestDto

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "JSON invalide")
		return
	}

	response, err := c.service.Register(data)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusCreated, response)
}
