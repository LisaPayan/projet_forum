package models

type Message struct {
	Id              int    `json:"id"`
	Contenu         string `json:"contenu"`
	PublishedAt     string `json:"date_publication"`
	User_c          User   `json:"user_c"`
	Fil_c           Fil    `json:"fil_c"`
	NbLikes         int    `json:"nb_likes"`
	NbDislikes      int    `json:"nb_dislikes"`
	ScorePopularite int    `json:"score_popularite"`
}
