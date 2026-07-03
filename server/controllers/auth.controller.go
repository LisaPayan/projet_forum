package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"projet_forum/auth"
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
	fmt.Println("test")
	contentType := r.Header.Get("Content-Type")

	if !strings.Contains(contentType, "application/json") {
		helper.WriteError(w, http.StatusBadRequest, "ereur format")
		return
	}

	var data dto.RegisterRequestDto
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "ereur decode")
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
	contentType := r.Header.Get("Content-Type")

	if !strings.Contains(contentType, "application/json") {
		helper.WriteError(w, http.StatusBadRequest, "ereur format")
		return
	}
	var data dto.LoginRequestDto
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "ereur decode")
		return
	}

	response, err := c.service.Login(data)
	if err != nil {
		helper.WriteError(w, http.StatusUnauthorized, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, response)
}

func (c *AuthController) Me(w http.ResponseWriter, r *http.Request) {
	claims, _ := r.Context().Value("user").(*auth.Claims)

	fmt.Println(claims)
	helper.WriteJSON(w, http.StatusOK, fmt.Sprintf("Hello %v , u are %v", claims.UserID, claims.Role))
}
