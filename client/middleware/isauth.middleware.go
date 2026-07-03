package middleware

import (
	"client/api"
	"context"
	"log"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, cookieErr := r.Cookie("access_token")
		if cookieErr != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		log.Printf("[MIDDLEWARE] Route appelée : %s", r.URL.Path)

		if r.URL.Path == "/dashboard" {
			// 🔴 AJOUTE /api À LA FIN DE L'URL ICI :
			authApi := api.InitAuthApi("http://localhost:8080/api")

			meData, err := authApi.Me(cookie.Value)

			if err != nil {
				log.Printf("🛑 [MIDDLEWARE] Erreur API Backend : %v", err)
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}

			log.Printf("🟢 [MIDDLEWARE] Réponse Backend -> Code: %d | Msg: '%s'", meData.Code, meData.Message)

			if !strings.Contains(strings.ToLower(meData.Message), "admin") {
				log.Printf("❌ [MIDDLEWARE] Accès refusé pour : '%s'", meData.Message)
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}

			log.Println("👑 [MIDDLEWARE] Accès autorisé au Dashboard !")
		}

		ctx := context.WithValue(r.Context(), "token", cookie.Value)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
