package middleware

import (
	"net/http"
	"strings"

	"projet_forum/server/auth"
	"projet_forum/server/helper"
)

func adminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			helper.WriteError(w, http.StatusUnauthorized, "token manquant")
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			helper.WriteError(w, http.StatusUnauthorized, "format de token invalide")
			return
		}

		claims, err := auth.ValidateToken(parts[1])
		if err != nil {
			helper.WriteError(w, http.StatusUnauthorized, "token invalide")
			return
		}

		if claims.Role != "admin" {
			helper.WriteError(w, http.StatusForbidden, "accès réservé aux administrateurs")
			return
		}

		next.ServeHTTP(w, r)
	})

}
