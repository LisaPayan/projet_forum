package repositories

import (
	"database/sql"
	"fmt"
	"projet_forum/models"
)

type AuthRepository struct {
	db *sql.DB
}

func InitAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) UserExists(pseudo string, email string) (bool, error) {
	var count int

	err := r.db.QueryRow(
		"SELECT COUNT(*) FROM users WHERE pseudo = ? OR email = ?",
		pseudo,
		email,
	).Scan(&count)

	if err != nil {
		return false, fmt.Errorf("erreur verification utilisateur : %s", err.Error())
	}

	return count > 0, nil
}

func (r *AuthRepository) Register(pseudo string, email string, password string) (int64, error) {
	result, err := r.db.Exec(
		"INSERT INTO users (pseudo, email, passwd) VALUES (?, ?, ?)",
		pseudo,
		email,
		password,
	)

	if err != nil {
		return 0, fmt.Errorf("erreur inscription utilisateur : %s", err.Error())
	}

	return result.LastInsertId()
}

func (r *AuthRepository) FindByUsernameOrEmail(username string) (*models.User, error) {
	var user models.User

	err := r.db.QueryRow(
		"SELECT id, pseudo, email, passwd, is_admin, is_ban FROM users WHERE pseudo = ? OR email = ?",
		username,
		username,
	).Scan(
		&user.Id,
		&user.Pseudo,
		&user.Email,
		&user.Passwd,
		&user.IsAdmin,
		&user.IsBan,
	)

	if err != nil {
		return nil, fmt.Errorf("utilisateur introuvable")
	}

	return &user, nil
}

func (r *AuthRepository) SaveToken(userID int, token string) error {
	_, err := r.db.Exec(
		"UPDATE users SET jwt = ? WHERE id = ?",
		token,
		userID,
	)

	if err != nil {
		return fmt.Errorf("erreur sauvegarde token : %s", err.Error())
	}

	return nil
}
