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
