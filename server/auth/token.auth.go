package auth

import (
	"errors"
	"fmt"
	"projet_forum/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userID string, role string) (string, error) {
	now := time.Now()
	secret := []byte(config.GetEnvWithDefault("JWT_SECRET", "secret_forum"))

	claims := Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			Issuer:    "forum-api",
			Audience:  []string{"forum-front"},
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(15 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}

// ValidateToken verifie un JWT et retourne ses informations si le token est valide.
func ValidateToken(tokenString string) (*Claims, error) {
	secret := []byte(config.GetEnvWithDefault("JWT_SECRET", "secret_forum"))
	// Claims recevra les donnees decodees depuis le token.
	claims := &Claims{}

	// ParseWithClaims analyse le token, verifie sa signature et remplit claims.
	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			// On refuse tout algorithme different de HS256 pour eviter les signatures inattendues.
			if token.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf(
					"unexpected signing method: %v",
					token.Header["alg"],
				)
			}
			// Cette cle doit etre la meme que celle utilisee pour signer le token.
			return secret, nil
		},
		// Ces options garantissent que le token vient de l'API attendue
		// et qu'il est destine au bon client.
		jwt.WithIssuer("forum-api"),
		jwt.WithAudience("forum-front"),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
	)

	// Une erreur ici signifie que le token est mal forme, expire ou invalide.
	if err != nil {
		return nil, err
	}

	// Par securite, on verifie aussi le statut final du token parse.
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Le token est valide : on retourne les claims utilisables par l'application.
	return claims, nil
}
