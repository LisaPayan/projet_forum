package dto

type FilDto struct {
	Id     int     `json:"id"`
	Titre  string  `json:"titre"`
	Statut string  `json:"statut"`
	Score  int     `json:"score"`
	User_c UserDto `json:"user_c"`
	Tag_c  TagDto  `json:"tag_c"`
}
