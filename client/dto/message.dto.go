package dto

type MessageDto struct {
	Id              int     `json:"id"`
	Contenu         string  `json:"contenu"`
	PublishedAt     string  `json:"date_publication"`
	User_c          UserDto `json:"user_c"`
	Fil_c           FilDto  `json:"fil_c"`
	NbLikes         int     `json:"nb_likes"`
	NbDislikes      int     `json:"nb_dislikes"`
	ScorePopularite int     `json:"score_popularite"`
}
