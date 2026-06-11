package models

type Message struct {
	Id          int    `json:"id"`
	Contenu     string `json:"contenu"`
	IsPublished int    `json:"is_published"`
	PublishedAt string `json:"date_publication"`
	NbrLikes    int    `json:"nbr_likes"`
	NbrDislikes int    `json:"nbr_dislikes"`
	User_c      User   `json:"user_c"`
	Fil
}
