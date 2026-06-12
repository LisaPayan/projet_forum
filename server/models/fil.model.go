package models

type Fil struct {
	Id     int    `json:"id"`
	Titre  string `json:"titre"`
	Statut string `json:"statut"`
	Score  int    `json:"score"`
	User_c User   `json:"user_c"`
	Tag_c  Tag    `json:"tag_c"`
}
