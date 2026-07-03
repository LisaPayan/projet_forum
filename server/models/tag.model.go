package models

type Tag struct {
	Id          int    `json:"id"`
	Nom         string `json:"nom"`
	Description string `json:"description"`
}
