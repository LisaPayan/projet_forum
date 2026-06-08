package dto

type FilDto struct {
	Id     int    `json:"id"`
	Titre  string `json:"titre"`
	Statut int    `json:"statut"`
	Score  int    `json:"score"`
}
