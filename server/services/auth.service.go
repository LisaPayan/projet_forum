package services

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"net/http"
	"projet_forum/dto"
	"projet_forum/repositories"
	"strings"
	"unicode"
)

type AuthService struct {
	authRepo *repositories.AuthRepository
}

func initAuthService(authRepo *repositories.AuthRepository) *AuthService {
	return	&AuthService{authRepo: authRepo}
}

func (s *AuthService) Register(request dto.RegisterRequestdto) (dto.RegisterResponsedto, error) {
	data.pseudo = strings.TrimSpace(data.pseudo)
	data.email = strings.TrimSpace(data.email)

	if data.pseudo == "" || data.email == "" || data.password == "" {
		return nil, errors.New("pseudo, email et mot de passe sont requis")
	}

if !validpassword(data.password) {
	return nil ,errors.New("mot de passe invalide")
}

exists, err := s.authRepo.userexists(data.pseudo, data.email)
if err != nil {
	return nil, err
}

if exists {
	return nil, errors.New("pseudo ou email déjà utilisé")
}

hashedPassword := hashPassword(data.password)

userId, err := s.authRepo.createUser(data.pseudo, data.email, hashedPassword)
if err != nil {
	return nil, err
}

return dto.RegisterResponsedto{
	code: http.StatusCreated,
	message: "Utilisateur inscrit avec succès",
	userId: userId,
}, nil
}

func validpassword(password string) bool {
	if len(password) <8 {
		return false
	}

hasupper := false
haspecial := false

for _, char := range password {
	if unicode.IsUpper(char) {
		hasupper = true
	}
}
	return hasupper && haspecial
}

func hashPassword(password string) string {
	hash := sha512.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}