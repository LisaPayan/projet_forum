package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"projet_forum/auth"
	"projet_forum/helper"
	"projet_forum/models"
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

func (c *FilControllers) CreateFil(w http.ResponseWriter, r *http.Request) {
	userData, _ := r.Context().Value("user").(*auth.Claims)
	var newFil models.Fil
	if err := json.NewDecoder(r.Body).Decode(&newFil); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "JSON invalide")
		return
	}

	userId, _ := strconv.Atoi(userData.UserID)
	newFil.User_c.Id = userId
	filId, filErr := c.service.Create(newFil)
	if filErr != nil {
		helper.WriteError(w, http.StatusBadRequest, filErr.Error())
		return
	}

	fil, filErr := c.service.ReadById(filId)
	if filErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, filErr.Error())
		return
	}

	helper.WriteJSON(w, http.StatusCreated, fil)
}

func (c *FilControllers) ReadById(w http.ResponseWriter, r *http.Request) {
	idFil, idFilErr := readFilId(r)
	if idFilErr != nil {
		helper.WriteError(w, http.StatusBadRequest, "Identifiant fil invalide")
		return
	}

	fil, filErr := c.service.ReadById(idFil)
	if filErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, filErr.Error())
		return
	}
	if fil.Id == 0 {
		helper.WriteError(w, http.StatusNotFound, "Produit introuvable")
		return
	}

	helper.WriteJSON(w, http.StatusOK, fil)
}

func (c *FilControllers) CreateMessage(w http.ResponseWriter, r *http.Request) {
	userData, _ := r.Context().Value("user").(*auth.Claims)

	vars := mux.Vars(r)
	filId, _ := strconv.Atoi(vars["id"])

	var newMessage models.Message
	if err := json.NewDecoder(r.Body).Decode(&newMessage); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "JSON invalide")
		return
	}

	userId, _ := strconv.Atoi(userData.UserID)

	newMessage.User_c.Id = userId
	newMessage.Fil_c.Id = filId

	messageId, messageErr := c.service.CreateMessage(newMessage)
	if messageErr != nil {
		helper.WriteError(w, http.StatusBadRequest, messageErr.Error())
		return
	}

	newMessage.Id = messageId

	helper.WriteJSON(w, http.StatusCreated, newMessage)
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

	tri := r.URL.Query().Get("tri")

	messagesList, filErr := c.service.FilByIdMessages(idFil, tri)
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

func (c *FilControllers) AjoutReaction(w http.ResponseWriter, r *http.Request) {
	userData, _ := r.Context().Value("user").(*auth.Claims)

	var newReaction models.Reaction
	if err := json.NewDecoder(r.Body).Decode(&newReaction); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "JSON invalide")
		return
	}

	userId, _ := strconv.Atoi(userData.UserID)
	newReaction.User_c.Id = userId
	// newReaction.Message_c.Id = messageId

	reac, reacErr := c.service.AjoutReaction(newReaction)
	if reacErr != nil {
		helper.WriteError(w, http.StatusBadRequest, reacErr.Error())
		return
	}

	// helper.WriteJSON(w, http.StatusCreated, fil)
	if reac == 0 {
		helper.WriteJSON(w, http.StatusOK, map[string]string{"message": "réaction annulée avec succès"})
	} else {
		helper.WriteJSON(w, http.StatusOK, map[string]string{"message": "réaction enregistrée avec succès"})
	}
}

func (c *FilControllers) UpdateFilById(w http.ResponseWriter, r *http.Request) {
	userData, _ := r.Context().Value("user").(*auth.Claims)

	idFil, idFilErr := readFilId(r)
	if idFilErr != nil {
		helper.WriteError(w, http.StatusBadRequest, "Identifiant fil invalide")
		return
	}

	userConnecteId, _ := strconv.Atoi(userData.UserID)

	idProprietaire, errOwner := c.service.GetFilOwner(idFil)
	if errOwner != nil {
		helper.WriteError(w, http.StatusNotFound, "Fil de discussion introuvable")
		return
	}

	if userConnecteId != idProprietaire {
		helper.WriteError(w, http.StatusForbidden, "Action refusée : Vous n'êtes pas le propriétaire de ce fil")
		return
	}

	var fil models.Fil
	if err := json.NewDecoder(r.Body).Decode(&fil); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "JSON invalide")
		return
	}
	fil.Id = idFil

	filErr := c.service.UpdateFilById(fil)
	if filErr != nil {
		helper.WriteError(w, http.StatusBadRequest, filErr.Error())
		return
	}

	updatedFil, filErr := c.service.ReadById(idFil)
	if filErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, filErr.Error())
		return
	}
	helper.WriteJSON(w, http.StatusOK, updatedFil)
}

func (c *FilControllers) DeleteFilById(w http.ResponseWriter, r *http.Request) {
	userData, _ := r.Context().Value("user").(*auth.Claims)

	idFil, idFilErr := readFilId(r)
	if idFilErr != nil {
		helper.WriteError(w, http.StatusBadRequest, "Identifiant fil invalide")
		return
	}

	userConnecteId, _ := strconv.Atoi(userData.UserID)
	isAdmin := userData.Role == "admin"

	idProprietaire, errOwner := c.service.GetFilOwner(idFil)
	if errOwner != nil {
		helper.WriteError(w, http.StatusNotFound, "Fil de discussion introuvable")
		return
	}

	if userConnecteId != idProprietaire && !isAdmin {
		helper.WriteError(w, http.StatusForbidden, "Action refusée : Vous n'avez pas les droits pour supprimer ce fil")
		return
	}

	filErr := c.service.DeleteFilById(idFil)
	if filErr != nil {
		helper.WriteError(w, http.StatusBadRequest, filErr.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "Fil supprime",
	})
}

func (c *FilControllers) UpdateMessageById(w http.ResponseWriter, r *http.Request) {
	userData, _ := r.Context().Value("user").(*auth.Claims)

	var message models.Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "JSON invalide")
		return
	}

	if message.Id == 0 {
		helper.WriteError(w, http.StatusBadRequest, "Identifiant du message manquant dans les données")
		return
	}

	userConnecteId, _ := strconv.Atoi(userData.UserID)

	idProprietaire, errOwner := c.service.GetMessageOwner(message.Id)
	if errOwner != nil {
		helper.WriteError(w, http.StatusNotFound, "Message introuvable")
		return
	}

	if userConnecteId != idProprietaire {
		helper.WriteError(w, http.StatusForbidden, "Action refusée : Vous n'êtes pas le propriétaire de ce message")
		return
	}

	messageErr := c.service.UpdateMessageById(message)
	if messageErr != nil {
		helper.WriteError(w, http.StatusBadRequest, messageErr.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "Message modifié avec succès",
	})
}
