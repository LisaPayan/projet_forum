package repositories

import (
	"database/sql"
	"fmt"
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
