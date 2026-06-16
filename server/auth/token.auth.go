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

func validatetoken(tokenString string) (*Claims, error) {
	secret := []byte(config.GetEnvWithDefault("JWT_SECRET", "secret_forum"))
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(t *jwt.Token) (interface{}, error) {
			if t == nil {
				return nil, fmt.Errorf("token is nil")
			}
			if t.Method == nil || t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, fmt.Errorf("invalid signing method")
			}
			return secret, nil
		},
		jwt.WithIssuer("forum-api"),
		jwt.WithAudience("forum-front"),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token invalide")
	}

	return claims, nil
}
