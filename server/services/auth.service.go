package services

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"net/http"
	"projet_forum/auth"
	"projet_forum/dto"
	"projet_forum/repositories"
	"strconv"
	"strings"
	"unicode"
)

type AuthService struct {
	authRepository *repositories.AuthRepository
}

func InitAuthService(authRepository *repositories.AuthRepository) *AuthService {
	return &AuthService{authRepository: authRepository}
}

func (s *AuthService) Register(data dto.RegisterRequestDto) (*dto.RegisterResponseDto, error) {
	data.Pseudo = strings.TrimSpace(data.Pseudo)
	data.Email = strings.TrimSpace(data.Email)

	if data.Pseudo == "" || data.Email == "" || data.Password == "" {
		return nil, errors.New("tous les champs sont obligatoires")
	}

	if !isPasswordValid(data.Password) {
		return nil, errors.New("le mot de passe doit contenir au minimum 12 caracteres, une majuscule et un caractere special")
	}

	exists, err := s.authRepository.UserExists(data.Pseudo, data.Email)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("pseudo ou email deja utilise")
	}

	hashedPassword := hashPassword(data.Password)

	userId, err := s.authRepository.Register(data.Pseudo, data.Email, hashedPassword)
	if err != nil {
		return nil, err
	}

	return &dto.RegisterResponseDto{
		Code:    http.StatusCreated,
		Message: "utilisateur inscrit avec succes",
		UserId:  userId,
	}, nil
}

func (s *AuthService) Login(data dto.LoginRequestDto) (*dto.LoginResponseDto, error) {
	data.Username = strings.TrimSpace(data.Username)

	if data.Username == "" || data.Password == "" {
		return nil, errors.New("identifiant et mot de passe obligatoires")
	}

	user, err := s.authRepository.FindByUsernameOrEmail(data.Username)
	if err != nil {
		return nil, errors.New("identifiants invalides")
	}

	if user.IsBan == 1 {
		return nil, errors.New("compte banni")
	}

	hashedPassword := hashPassword(data.Password)
	if hashedPassword != user.Passwd {
		return nil, errors.New("identifiants invalides")
	}

	role := "user"
	if user.IsAdmin == 1 {
		role = "admin"
	}

	token, err := auth.GenerateToken(strconv.Itoa(user.Id), role)
	if err != nil {
		return nil, err
	}

	err = s.authRepository.SaveToken(user.Id, token)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponseDto{
		Type:        "Bearer",
		AccessToken: token,
		ExpiresIn:   900,
	}, nil
}

func isPasswordValid(password string) bool {
	if len(password) < 12 {
		return false
	}

	hasUpper := false
	hasSpecial := false

	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpper = true
		}

		if unicode.IsPunct(char) || unicode.IsSymbol(char) {
			hasSpecial = true
		}
	}

	return hasUpper && hasSpecial
}

func hashPassword(password string) string {
	hash := sha512.Sum512([]byte(password))
	return hex.EncodeToString(hash[:])
}
