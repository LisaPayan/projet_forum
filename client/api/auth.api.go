package api

type AuthApi struct {
	BaseURL string
}

func InitAuthApi(baseURL string) *AuthApi {
	return &AuthApi{BaseURL: baseURL}
}
