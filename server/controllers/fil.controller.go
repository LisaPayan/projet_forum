package controllers

import (
	"fmt"
	"net/http"
	"projet_forum/helper"
	"projet_forum/services"
	"strconv"

	"github.com/gorilla/mux"
)

type FilControllers struct {
	service *services.FilService
}

func InitFilController(service *services.FilService) *FilControllers {
	return &FilControllers{service: service}
}

func readFilId(r *http.Request) (int, error) {
	return strconv.Atoi(mux.Vars(r)["id"])
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

func (c *FilControllers) GetMessagesByFil(w http.ResponseWriter, r *http.Request) {
	idFil, idFilErr := readFilId(r)
	if idFilErr != nil {
		helper.WriteError(w, http.StatusBadRequest, "Identifiant fil invalide")
		return
	}

	messagesList, filErr := c.service.FilByIdMessages(idFil)
	if filErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, filErr.Error())
		return
	}

	if len(messagesList) == 0 {
		helper.WriteError(w, http.StatusNotFound, "Aucun message trouvé pour ce fil")
		return
	}

	helper.WriteJSON(w, http.StatusOK, messagesList)
	fmt.Println(messagesList)

}

func (c *FilControllers) GetMessagesByFilAnciens(w http.ResponseWriter, r *http.Request) {
	idFil, idFilErr := readFilId(r)
	if idFilErr != nil {
		helper.WriteError(w, http.StatusBadRequest, "Identifiant fil invalide")
		return
	}

	messagesList, filErr := c.service.FilByIdMessagesAnciens(idFil)
	if filErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, filErr.Error())
		return
	}

	if len(messagesList) == 0 {
		helper.WriteError(w, http.StatusNotFound, "Aucun message trouvé pour ce fil")
		return
	}

	helper.WriteJSON(w, http.StatusOK, messagesList)
	fmt.Println(messagesList)

}

func (c *FilControllers) GetMessagesByFilRecents(w http.ResponseWriter, r *http.Request) {
	idFil, idFilErr := readFilId(r)
	if idFilErr != nil {
		helper.WriteError(w, http.StatusBadRequest, "Identifiant fil invalide")
		return
	}

	messagesList, filErr := c.service.FilByIdMessagesRecents(idFil)
	if filErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, filErr.Error())
		return
	}

	if len(messagesList) == 0 {
		helper.WriteError(w, http.StatusNotFound, "Aucun message trouvé pour ce fil")
		return
	}

	helper.WriteJSON(w, http.StatusOK, messagesList)
	fmt.Println(messagesList)

}

func (c *FilControllers) FilsPetanque(w http.ResponseWriter, r *http.Request) {
	filsList, filsErr := c.service.FilsPetanque()
	if filsErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, filsErr.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, filsList)
	fmt.Println(filsList)
}

func (c *FilControllers) FilsCuisine(w http.ResponseWriter, r *http.Request) {
	filsList, filsErr := c.service.FilsCuisine()
	if filsErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, filsErr.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, filsList)
	fmt.Println(filsList)
}

func (c *FilControllers) FilsNature(w http.ResponseWriter, r *http.Request) {
	filsList, filsErr := c.service.FilsNature()
	if filsErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, filsErr.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, filsList)
	fmt.Println(filsList)
}
