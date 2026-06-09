package dto

type RegisterRequest struct {
	code int	`json:"code"`
	message string `json:"message"`
	userId int `json:"userId"`
}