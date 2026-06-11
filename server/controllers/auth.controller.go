package controllers*

import (
	"net/http"
	"projet_forum/dto"
	"projet_forum/services"
	"encoding/json"
	"projet_forum/helpers"
)

type AuthController struct {
	service *services.AuthService
}

func initAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{service: authService}
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var data dto.RegisterRequestdto

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		helpers.writeerrorResponse(w, http.StatusBadRequest, "JSON invalide")
		return
	}

	response, err := c.service.Register(data)
	if err != nil {
		helpers.writeerror(w, http.StatusBadRequest, err.Error()
		return
	}

	helpers.writeJSONResponse(w, http.StatusCreated, response)
}