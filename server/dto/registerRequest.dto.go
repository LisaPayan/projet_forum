package dto

type RegisterRequest struct {
	pseudo string 'json:"pseudo" '
	email  string 'json:"email" '
	password string 'json:"password" '
}