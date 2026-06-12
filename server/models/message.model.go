package models

type Message struct {
	Id          int    `json:"id"`
	Contenu     string `json:"contenu"`
	PublishedAt string `json:"date_publication"`
	User_c      User   `json:"user_c"`
	Fil
}
