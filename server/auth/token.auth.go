package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// jwtSecret est la cle de signature des JWT, chargee une seule fois au demarrage via Init.
var jwtSecret []byte

// Init enregistre la cle de signature des JWT. Elle doit etre appelee au demarrage.
func Init(secret string) {
	jwtSecret = []byte(secret)
}

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	IsBan  int    `json:"is_ban"`
	jwt.RegisteredClaims
}

func GenerateToken(userID string, role string, isBan int) (string, error) {
	now := time.Now()
	if len(jwtSecret) == 0 {
		return "", errors.New("JWT secret non initialise")
	}

	claims := Claims{
		UserID: userID,
		Role:   role,
		IsBan:  isBan,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			Issuer:    "forum-api",
			Audience:  []string{"forum-front"},
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(15 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

// ValidateToken verifie un JWT et retourne ses informations si le token est valide.
func ValidateToken(tokenString string) (*Claims, error) {
	if len(jwtSecret) == 0 {
		return nil, errors.New("JWT secret non initialise")
	}
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
			return jwtSecret, nil
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

func IsAdmin(claims *Claims) bool {
	if claims == nil {
		return false
	}
	return claims.Role == "admin"
}
