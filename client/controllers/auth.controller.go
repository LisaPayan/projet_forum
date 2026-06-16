package controllers

import (
	"client/services"
	"client/templates"
	"encoding/json"
	"net/http"
)

// AuthControllers gere les pages liees a l'authentification.
type AuthControllers struct {
	service  *services.AuthService
	template *templates.TemplateManager
}

// InitAuthController cree un controller auth avec son service et son moteur de templates.
func InitAuthController(service *services.AuthService, template *templates.TemplateManager) *AuthControllers {
	return &AuthControllers{service: service, template: template}
}

// LoginForm affiche la page de connexion.
func (c *AuthControllers) LoginForm(w http.ResponseWriter, r *http.Request) {
	c.template.RenderTemplate(w, r, "login", nil)
}

// Login traite la soumission du formulaire de connexion.
func (c *AuthControllers) Login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		c.template.RenderTemplate(w, r, "login", "Formulaire invalide")
		return
	}

	// On recupere les champs HTML puis on delegue la verification au service.
	token, err := c.service.Login(r.FormValue("username"), r.FormValue("password"))
	if err != nil {
		c.template.RenderTemplate(w, r, "login", err.Error())
		return
	}

	// Le token est conserve dans un cookie pour pouvoir le reutiliser sur les prochaines pages.
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   24 * 60 * 60,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Logout supprime le cookie contenant le token puis redirige vers l'accueil.
func (c *AuthControllers) Logout(w http.ResponseWriter, r *http.Request) {

	// MaxAge a -1 indique au navigateur de supprimer le cookie.
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Me affiche la reponse de la route /me pour l'utilisateur connecte.
// Le controller lit le token dans le cookie, appelle le service, puis renvoie le DTO en JSON.
func (c *AuthControllers) Me(w http.ResponseWriter, r *http.Request) {
	// Sans cookie access_token, l'utilisateur n'est pas authentifie sur le client.
	cookies, cookieErr := r.Cookie("access_token")
	if cookieErr != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	value, err := c.service.Me(cookies.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Indique au navigateur que la reponse est du JSON.
	w.Header().Set("Content-Type", "application/json")
	// Encode le DTO MeResponseDto pour produire :
	// {"code":200,"message":"Hello user 1 with role admin"}
	json.NewEncoder(w).Encode(value)
}

func (c *AuthControllers) RegisterForm(w http.ResponseWriter, r *http.Request) {
	c.template.RenderTemplate(w, r, "register", nil)
}

func (c *AuthControllers) Register(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		c.template.RenderTemplate(w, r, "register", "Formulaire invalide")
		return
	}

	pseudo := r.FormValue("pseudo")
	email := r.FormValue("email")
	password := r.FormValue("password")

	_, err := c.service.Register(pseudo, email, password)
	if err != nil {
		c.template.RenderTemplate(w, r, "register", err.Error())
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
