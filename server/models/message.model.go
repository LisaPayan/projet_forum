package models

type Message struct {
	Id          int    `json:"id"`
	IsPublished int    `json:"is_published"`
	PublishedAt string `json:"date_publication"`
	NbrLikes    int    `json:"nbr_likes"`
	NbrDislikes int    `json:"nbr_dislikes"`
}
