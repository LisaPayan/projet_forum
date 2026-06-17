package controllers

import (
	"client/dto"
	"client/services"
	"client/templates"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"unicode"

	"github.com/gorilla/mux"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
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

// func (c *FilControllers) DisplayMessagesFil(w http.ResponseWriter, r *http.Request) {
// 	idFil, idFilErr := strconv.Atoi(mux.Vars(r)["id"])
// 	if idFilErr != nil {
// 		http.Error(w, "Erreur - Identifiant produit invalide", http.StatusBadRequest)
// 		return
// 	}

// 	fil, filErr := c.service.ReadByIdMessages(idFil)
// 	if filErr != nil {
// 		http.Error(w, filErr.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	c.template.RenderTemplate(w, r, "details_fil", fil)
// }

// func (c *FilControllers) DisplayListPetanque(w http.ResponseWriter, r *http.Request) {
// 	filList, filErr := c.service.FilsPetanque()
// 	if filErr != nil {
// 		http.Error(w, filErr.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	c.template.RenderTemplate(w, r, "petanque", filList)
// }

// func (c *FilControllers) DisplayListCuisine(w http.ResponseWriter, r *http.Request) {
// 	filList, filErr := c.service.FilsCuisine()
// 	if filErr != nil {
// 		http.Error(w, filErr.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	c.template.RenderTemplate(w, r, "cuisine", filList)
// }

// func (c *FilControllers) DisplayListNature(w http.ResponseWriter, r *http.Request) {
// 	filList, filErr := c.service.FilsNature()
// 	if filErr != nil {
// 		http.Error(w, filErr.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	c.template.RenderTemplate(w, r, "nature", filList)
// }

func (c *FilControllers) CreateForm(w http.ResponseWriter, r *http.Request) {
	c.template.RenderTemplate(w, r, "fil_create", nil)
}

func (c *FilControllers) Create(w http.ResponseWriter, r *http.Request) {

	tagIdStr := r.FormValue("tag_id")
	tagIdInt, _ := strconv.Atoi(tagIdStr)

	newFil := dto.FilDto{}
	newFil.Titre = r.FormValue("titre")
	newFil.Tag_c.Id = tagIdInt

	cookie, err := r.Cookie("access_token")
	if err != nil || cookie == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	filId, err := c.service.Create(newFil, cookie.Value)
	if err != nil {
		referer := r.Referer()
		if referer == "" {
			referer = "/all"
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	destination := fmt.Sprintf("/fil/%d/messages", filId)
	http.Redirect(w, r, destination, http.StatusSeeOther)
}

func (c *FilControllers) CreateMessage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	filId, _ := strconv.Atoi(vars["id"])

	newMessage := dto.MessageDto{}
	newMessage.Contenu = r.FormValue("contenu")

	cookie, err := r.Cookie("access_token")
	if err != nil || cookie == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	_, err = c.service.CreateMessage(newMessage, filId, cookie.Value)
	if err != nil {
		referer := r.Referer()
		if referer == "" {
			referer = fmt.Sprintf("/fil/%d/messages", filId)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	destination := fmt.Sprintf("/fil/%d/messages", filId)
	http.Redirect(w, r, destination, http.StatusSeeOther)
}

func (c *FilControllers) DisplayPaginationAll(w http.ResponseWriter, r *http.Request) {
	pageStr := r.FormValue("page")
	pageInt, _ := strconv.Atoi(pageStr)

	if pageInt < 0 {
		pageInt = 0
	}

	nbr_vis := r.FormValue("nbr_vis")
	if nbr_vis == "" {
		nbr_vis = "10"
	}
	data, err := c.service.ReadAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nbrInt := 10
	if nbr_vis == "all" {
		nbrInt = len(data)
	} else if nbr_vis == "20" {
		nbrInt = 20
	} else if nbr_vis == "30" {
		nbrInt = 30
	} else {
		nbrInt = 10
	}

	startIndex := pageInt * nbrInt
	endIndex := startIndex + nbrInt

	if startIndex >= len(data) {
		pageInt = 0
		startIndex = 0
		endIndex = nbrInt
	}

	if endIndex > len(data) {
		endIndex = len(data)
	}

	SelectFils := data[startIndex:endIndex]

	prevPage := pageInt - 1
	if prevPage < 0 {
		prevPage = 0
	}

	nextPage := pageInt
	if endIndex < len(data) {
		nextPage = pageInt + 1
	}

	vieData := dto.PagePaginationFil{
		Page:   pageInt,
		Next:   nextPage,
		Prev:   prevPage,
		NbrVis: nbr_vis,
		Data:   SelectFils,
	}

	c.template.RenderTemplate(w, r, "all", vieData)

}

func (c *FilControllers) DisplayPaginationPetanque(w http.ResponseWriter, r *http.Request) {
	pageStr := r.FormValue("page")
	pageInt, _ := strconv.Atoi(pageStr)

	if pageInt < 0 {
		pageInt = 0
	}

	nbr_vis := r.FormValue("nbr_vis")
	if nbr_vis == "" {
		nbr_vis = "10"
	}
	data, err := c.service.FilsPetanque()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nbrInt := 10
	if nbr_vis == "all" {
		nbrInt = len(data)
	} else if nbr_vis == "20" {
		nbrInt = 20
	} else if nbr_vis == "30" {
		nbrInt = 30
	} else {
		nbrInt = 10
	}

	startIndex := pageInt * nbrInt
	endIndex := startIndex + nbrInt

	if startIndex >= len(data) {
		pageInt = 0
		startIndex = 0
		endIndex = nbrInt
	}

	if endIndex > len(data) {
		endIndex = len(data)
	}

	SelectFils := data[startIndex:endIndex]

	prevPage := pageInt - 1
	if prevPage < 0 {
		prevPage = 0
	}

	nextPage := pageInt
	if endIndex < len(data) {
		nextPage = pageInt + 1
	}

	vieData := dto.PagePaginationFil{
		Page:   pageInt,
		Next:   nextPage,
		Prev:   prevPage,
		NbrVis: nbr_vis,
		Data:   SelectFils,
	}

	c.template.RenderTemplate(w, r, "petanque", vieData)

}

func (c *FilControllers) DisplayPaginationCuisine(w http.ResponseWriter, r *http.Request) {
	pageStr := r.FormValue("page")
	pageInt, _ := strconv.Atoi(pageStr)

	if pageInt < 0 {
		pageInt = 0
	}

	nbr_vis := r.FormValue("nbr_vis")
	if nbr_vis == "" {
		nbr_vis = "10"
	}
	data, err := c.service.FilsCuisine()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nbrInt := 10
	if nbr_vis == "all" {
		nbrInt = len(data)
	} else if nbr_vis == "20" {
		nbrInt = 20
	} else if nbr_vis == "30" {
		nbrInt = 30
	} else {
		nbrInt = 10
	}

	startIndex := pageInt * nbrInt
	endIndex := startIndex + nbrInt

	if startIndex >= len(data) {
		pageInt = 0
		startIndex = 0
		endIndex = nbrInt
	}

	if endIndex > len(data) {
		endIndex = len(data)
	}

	SelectFils := data[startIndex:endIndex]

	prevPage := pageInt - 1
	if prevPage < 0 {
		prevPage = 0
	}

	nextPage := pageInt
	if endIndex < len(data) {
		nextPage = pageInt + 1
	}

	vieData := dto.PagePaginationFil{
		Page:   pageInt,
		Next:   nextPage,
		Prev:   prevPage,
		NbrVis: nbr_vis,
		Data:   SelectFils,
	}

	c.template.RenderTemplate(w, r, "cuisine", vieData)

}

func (c *FilControllers) DisplayPaginationNature(w http.ResponseWriter, r *http.Request) {
	pageStr := r.FormValue("page")
	pageInt, _ := strconv.Atoi(pageStr)

	if pageInt < 0 {
		pageInt = 0
	}

	nbr_vis := r.FormValue("nbr_vis")
	if nbr_vis == "" {
		nbr_vis = "10"
	}
	data, err := c.service.FilsNature()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nbrInt := 10
	if nbr_vis == "all" {
		nbrInt = len(data)
	} else if nbr_vis == "20" {
		nbrInt = 20
	} else if nbr_vis == "30" {
		nbrInt = 30
	} else {
		nbrInt = 10
	}

	startIndex := pageInt * nbrInt
	endIndex := startIndex + nbrInt

	if startIndex >= len(data) {
		pageInt = 0
		startIndex = 0
		endIndex = nbrInt
	}

	if endIndex > len(data) {
		endIndex = len(data)
	}

	SelectFils := data[startIndex:endIndex]

	prevPage := pageInt - 1
	if prevPage < 0 {
		prevPage = 0
	}

	nextPage := pageInt
	if endIndex < len(data) {
		nextPage = pageInt + 1
	}

	vieData := dto.PagePaginationFil{
		Page:   pageInt,
		Next:   nextPage,
		Prev:   prevPage,
		NbrVis: nbr_vis,
		Data:   SelectFils,
	}

	c.template.RenderTemplate(w, r, "nature", vieData)

}

func (c *FilControllers) DisplayPaginationMessage(w http.ResponseWriter, r *http.Request) {
	idFil, idFilErr := strconv.Atoi(mux.Vars(r)["id"])
	if idFilErr != nil {
		http.Error(w, "Erreur - Identifiant produit invalide", http.StatusBadRequest)
		return
	}
	pageStr := r.FormValue("page")
	pageInt, _ := strconv.Atoi(pageStr)

	if pageInt < 0 {
		pageInt = 0
	}

	nbr_vis := r.FormValue("nbr_vis")
	if nbr_vis == "" {
		nbr_vis = "10"
	}

	tri := r.FormValue("tri")

	data, err := c.service.ReadByIdMessages(idFil, tri)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nbrInt := 10
	if nbr_vis == "all" {
		nbrInt = len(data)
	} else if nbr_vis == "20" {
		nbrInt = 20
	} else if nbr_vis == "30" {
		nbrInt = 30
	} else {
		nbrInt = 10
	}

	startIndex := pageInt * nbrInt
	endIndex := startIndex + nbrInt

	if startIndex >= len(data) {
		pageInt = 0
		startIndex = 0
		endIndex = nbrInt
	}

	if endIndex > len(data) {
		endIndex = len(data)
	}

	SelectMessages := data[startIndex:endIndex]

	prevPage := pageInt - 1
	if prevPage < 0 {
		prevPage = 0
	}

	nextPage := pageInt
	if endIndex < len(data) {
		nextPage = pageInt + 1
	}

	vieData := dto.PagePaginationMess{
		IdFil:  idFil,
		Page:   pageInt,
		Next:   nextPage,
		Prev:   prevPage,
		NbrVis: nbr_vis,
		Data:   SelectMessages,
		Tri:    tri,
	}

	c.template.RenderTemplate(w, r, "details_fil", vieData)
}

func removeAccentsAndLower(str string) string {
	str = strings.ToLower(str)

	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, str)

	return result
}

func (c *FilControllers) DisplaySearch(w http.ResponseWriter, r *http.Request) {
	queryRaw := r.FormValue("query")
	queryClean := strings.TrimSpace(removeAccentsAndLower(queryRaw))

	if queryClean == "" {
		http.Redirect(w, r, "/search", http.StatusSeeOther)
		return
	}

	pageStr := r.FormValue("page")
	pageInt, _ := strconv.Atoi(pageStr)
	if pageInt < 0 {
		pageInt = 0
	}

	nbr_vis := r.FormValue("nbr_vis")
	if nbr_vis == "" {
		nbr_vis = "10"
	}

	data, dataError := c.service.ReadAll()
	if dataError != nil {
		http.Error(w, dataError.Error(), http.StatusInternalServerError)
		return
	}

	searchList := []dto.FilDto{}
	for _, item := range data {
		checkTitre := strings.Contains(removeAccentsAndLower(item.Titre), queryClean)
		checkUser := strings.Contains(removeAccentsAndLower(item.User_c.Pseudo), queryClean)
		checkTag := strings.Contains(removeAccentsAndLower(item.Tag_c.Nom), queryClean)

		if checkTitre || checkUser || checkTag {
			searchList = append(searchList, item)
		}
	}

	nbrInt := 10
	if nbr_vis == "all" {
		nbrInt = len(searchList)
	} else if nbr_vis == "20" {
		nbrInt = 20
	} else if nbr_vis == "30" {
		nbrInt = 30
	}

	totalItems := len(searchList)
	var SelectFils []dto.FilDto
	startIndex := pageInt * nbrInt

	if totalItems == 0 {
		pageInt = 0
		startIndex = 0
		SelectFils = []dto.FilDto{}
	} else {
		if startIndex >= totalItems {
			pageInt = (totalItems - 1) / nbrInt
			startIndex = pageInt * nbrInt
		}

		endIndex := startIndex + nbrInt
		if endIndex > totalItems {
			endIndex = totalItems
		}
		SelectFils = searchList[startIndex:endIndex]
	}

	prevPage := pageInt - 1
	if prevPage < 0 {
		prevPage = 0
	}

	nextPage := pageInt
	if startIndex+nbrInt < totalItems {
		nextPage = pageInt + 1
	}

	vieData := dto.PagePaginationFil{
		Query:  queryRaw,
		Page:   pageInt,
		Next:   nextPage,
		Prev:   prevPage,
		NbrVis: nbr_vis,
		Data:   SelectFils,
	}

	c.template.RenderTemplate(w, r, "search", vieData)
}

func (c *FilControllers) AjoutReaction(w http.ResponseWriter, r *http.Request) {

	messageIdStr := r.FormValue("message_id")
	messagedInt, _ := strconv.Atoi(messageIdStr)

	typeReaction := r.FormValue("type_reaction")
	filIdStr := r.FormValue("fil_id")

	newReac := dto.Reaction{}
	newReac.Message_c.Id = messagedInt
	newReac.Type_reac = typeReaction

	cookie, err := r.Cookie("access_token")
	if err != nil || cookie == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	_, err = c.service.AjoutReaction(newReac, cookie.Value)
	if err != nil {
		log.Println(" ERREUR REACTION CONTROLLER :", err)
		destination := "/fil/" + filIdStr + "/messages"
		http.Redirect(w, r, destination, http.StatusSeeOther)
		return
	}

	destination := "/fil/" + filIdStr + "/messages"
	http.Redirect(w, r, destination, http.StatusSeeOther)
}

func (c *FilControllers) UpdateForm(w http.ResponseWriter, r *http.Request) {
	idFil, _ := strconv.Atoi(mux.Vars(r)["id"])

	filActuel, err := c.service.ReadById(idFil)
	if err != nil {
		http.Error(w, "Impossible de charger le fil de discussion", http.StatusInternalServerError)
		return
	}

	c.template.RenderTemplate(w, r, "fil_update", map[string]interface{}{
		"Fil": filActuel,
	})
}

func (c *FilControllers) UpdateFilById(w http.ResponseWriter, r *http.Request) {
	idFil, _ := strconv.Atoi(mux.Vars(r)["id"])
	tagIdStr := r.FormValue("tag_id")
	tagIdInt, _ := strconv.Atoi(tagIdStr)

	updateFil := dto.FilDto{}
	updateFil.Id = idFil
	updateFil.Titre = r.FormValue("titre")
	updateFil.Tag_c.Id = tagIdInt

	cookie, err := r.Cookie("access_token")
	if err != nil || cookie == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	err = c.service.UpdateFilById(updateFil, cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/fil/%d/messages", idFil), http.StatusSeeOther)
}

func (c *FilControllers) DeleteFilById(w http.ResponseWriter, r *http.Request) {
	idFil, _ := strconv.Atoi(mux.Vars(r)["id"])

	cookie, err := r.Cookie("access_token")
	if err != nil || cookie == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	err = c.service.DeleteFilById(idFil, cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/all", http.StatusSeeOther)
}

func (c *FilControllers) UpdateFormMessage(w http.ResponseWriter, r *http.Request) {
	idMessageStr := r.URL.Query().Get("message_id")
	idFilStr := r.URL.Query().Get("fil_id")
	contenuActuel := r.URL.Query().Get("contenu")

	c.template.RenderTemplate(w, r, "message_update", map[string]interface{}{
		"MessageId":     idMessageStr,
		"IdFil":         idFilStr,
		"ContenuActuel": contenuActuel,
	})
}

func (c *FilControllers) UpdateMessageById(w http.ResponseWriter, r *http.Request) {
	idMessageStr := r.FormValue("message_id")
	idMessage, _ := strconv.Atoi(idMessageStr)

	nouveauContenu := r.FormValue("contenu")
	idFilStr := r.FormValue("fil_id")

	cookie, err := r.Cookie("access_token")
	if err != nil || cookie == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	updateMessage := dto.MessageDto{
		Id:      idMessage,
		Contenu: nouveauContenu,
	}

	err = c.service.UpdateMessageById(updateMessage, cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/fil/%s/messages", idFilStr), http.StatusSeeOther)
}

func (c *FilControllers) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	idMessage, _ := strconv.Atoi(r.FormValue("message_id"))
	idFilStr := r.FormValue("fil_id")

	cookie, err := r.Cookie("access_token")
	if err != nil || cookie == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	err = c.service.DeleteMessage(idMessage, cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/fil/%s/messages", idFilStr), http.StatusSeeOther)
}
