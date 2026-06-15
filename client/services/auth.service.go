package services

import (
	"client/api"
)

type AuthService struct {
	authApi *api.AuthApi
}

func InitAuthService(authApi *api.AuthApi) *AuthService {
	return &AuthService{authApi: authApi}
}
