package controllers

import (
	"fmt"
	"net/http"
	"projet_forum/helper"
	"projet_forum/services"
)

type FilControllers struct {
	service *services.FilService
}

func InitFilController(service *services.FilService) *FilControllers {
	return &FilControllers{service: service}
}

func (c *FilControllers) ReadAll(w http.ResponseWriter, r *http.Request) {
	filsList, filsErr := c.service.ReadAll()
	if filsErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, filsErr.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, filsList)
	fmt.Println(filsList)
}
