package middleware

import (
	"net/http"
	"strings"
)

func adminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			helper.writeerror(w, http.StatusUnauthorized, "token manquant")
			return
		}

		parts := strings.SplitN(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			helper.writeerror(w, http.StatusUnauthorized, "format de token invalide")
			return
		}

		claims, err := auth.validatetoken(parts[1])
		if err != nil {
			helper.writeerror(w, http.StatusUnauthorized, "token invalide")
			return
		}

		if claims.Role != "admin" {
			helper.writeerror(w, http.StatusForbidden, "accès réservé aux administrateurs")
			return
		}

		next.ServeHTTP(w, r)
	})

}
