package dto

type RegisterResponseDto struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	UserId  int64  `json:"user_id"`
}
