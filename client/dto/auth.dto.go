package dto

type LoginRequestDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponseDto struct {
	Type        string `json:"type"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type MeResponseDto struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type RegisterRequestDto struct {
	Pseudo   string `json:"pseudo"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponseDto struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	UserId  int64  `json:"user_id"`
}

type UserResponseDto struct {
	Id      int64  `json:"id"`
	Pseudo  string `json:"pseudo"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	IsAdmin bool   `json:"is_admin"`
}
