package dto

type UserResponseDto struct {
	Id      int64  `json:"id"`
	Pseudo  string `json:"pseudo"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	IsAdmin bool   `json:"is_admin"`
}
