package services

import (
	"client/api"
	"client/dto"
	"fmt"
)

type AuthService struct {
	authApi *api.AuthApi
}

// InitAuthService initialise le service d'authentification avec son client API.
func InitAuthService(authApi *api.AuthApi) *AuthService {
	return &AuthService{authApi: authApi}
}

// Login verifie les donnees du formulaire puis delegue l'appel HTTP au client API.
func (s *AuthService) Login(username string, password string) (string, error) {
	if username == "" || password == "" {
		return "", fmt.Errorf("Identifiant et mot de passe obligatoires")
	}

	// Le service garde les validations metier, mais ne construit plus la requete HTTP.
	response, err := s.authApi.Login(dto.LoginRequestDto{
		Username: username,
		Password: password,
	})
	if err != nil {
		return "", err
	}

	if response.AccessToken == "" {
		return "", fmt.Errorf("Token manquant dans la reponse")
	}

	return response.AccessToken, nil
}

func (s *AuthService) Me(token string) (dto.MeResponseDto, error) {
	result, resultErr := s.authApi.Me(token)
	if resultErr != nil {
		return dto.MeResponseDto{}, fmt.Errorf("Erreur service (auth.me) - Une erreur s'est produite : %v", resultErr)
	}
	return result, nil
}

func (s *AuthService) Register(pseudo, email, password string) (dto.RegisterResponseDto, error) {
	if pseudo == "" || email == "" || password == "" {
		return dto.RegisterResponseDto{}, fmt.Errorf("Tous les champs sont obligatoires")
	}

	response, err := s.authApi.Register(dto.RegisterRequestDto{
		Pseudo:   pseudo,
		Email:    email,
		Password: password,
	})
	if err != nil {
		return dto.RegisterResponseDto{}, err
	}

	return response, nil
}
