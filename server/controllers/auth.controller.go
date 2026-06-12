package controllers

import (
	"encoding/json"
	"net/http"
	"projet_forum/dto"
	"projet_forum/helper"
	"projet_forum/services"
	"strings"
)

type AuthController struct {
	service *services.AuthService
}

func InitAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{service: authService}
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var data dto.RegisterRequestDto

	if err := readRegisterRequest(r, &data); err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	response, err := c.service.Register(data)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusCreated, response)
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var data dto.LoginRequestDto

	if err := readLoginRequest(r, &data); err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	response, err := c.service.Login(data)
	if err != nil {
		helper.WriteError(w, http.StatusUnauthorized, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, response)
}

func readRegisterRequest(r *http.Request, data *dto.RegisterRequestDto) error {
	contentType := r.Header.Get("Content-Type")

	if strings.Contains(contentType, "application/json") {
		return json.NewDecoder(r.Body).Decode(data)
	}

	if err := r.ParseForm(); err != nil {
		return err
	}

	data.Pseudo = r.FormValue("pseudo")
	data.Email = r.FormValue("email")
	data.Password = r.FormValue("password")

	return nil
}

func readLoginRequest(r *http.Request, data *dto.LoginRequestDto) error {
	contentType := r.Header.Get("Content-Type")

	if strings.Contains(contentType, "application/json") {
		return json.NewDecoder(r.Body).Decode(data)
	}

	if err := r.ParseForm(); err != nil {
		return err
	}

	data.Username = r.FormValue("username")
	data.Password = r.FormValue("password")

	return nil
}
