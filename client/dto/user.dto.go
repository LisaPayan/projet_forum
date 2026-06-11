package dto

type UserDto struct {
	Id       int    `json:"id"`
	Pseudo   string `json:"pseudo"`
	Mail     string `json:"email"`
	Passwd   string `json:"passwd"`
	CreateAt string `json:"date_inscription"`
	IsAdmin  int    `json:"is_admin"`
	JWT      string `json:"JWT"`
	IsBan    int    `json:"is_ban"`
}
