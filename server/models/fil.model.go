package models

type Fil struct {
	Id     int    `json:"id"`
	Titre  string `json:"titre"`
	Statut int    `json:"statut"`
	Score  int    `json:"score"`
	User
}
