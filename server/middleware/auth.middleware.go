package middleware

import (
	"context"
	"fmt"
	"net/http"
	"projet_forum/auth"
	"projet_forum/helper"
	"strings"
)

// AuthMiddleware protege une route en exigeant un JWT valide.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Le token doit etre envoye dans le header Authorization.
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			helper.WriteError(w, http.StatusUnauthorized, "missing authorization header")
			return
		}

		// Le format attendu est : Bearer <token>.
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			helper.WriteError(w, http.StatusUnauthorized, "invalid authorization header")
			return
		}

		// On valide le token et on recupere les claims de l'utilisateur.
		claims, err := auth.ValidateToken(parts[1])
		if err != nil {
			fmt.Println(err)
			helper.WriteError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		if claims.IsBan == 1 {
			helper.WriteError(w, http.StatusForbidden, "votre compte a été banni de la plateforme")
			return
		}

		if strings.HasPrefix(r.URL.Path, "/dashboard") && strings.ToLower(claims.Role) != "admin" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		// Les claims sont ajoutes au contexte pour etre reutilises par les handlers suivants.
		ctx := context.WithValue(r.Context(), "user", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
