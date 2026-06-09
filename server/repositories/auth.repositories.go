package repositories

import (
	"database/sql"
	"fmt"
)

type AuthRepository struct {
	db *sql.DB
}

func initAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) userexists(pseudo string, email string) (bool, error) {
	var count int

	err := r.db.QueryRow(
		select count(*) from users where pseudo = ? or email = ?,
		pseudo,
		email,
	).Scan(&count)

	if err != nil {
		return false, fmt.Errorf("failed to check if user exists: %w", err)
	}

	return count > 0, nil
}

func(r *AuthRepository) createUser(pseudo string, email string, password string) (int64, error) {
	result, err := r.db.Exec(
	"insert into users (pseudo, email, password) values (?, ?, ?)",
	pseudo,
	email,
	password,
	)

	if err != nil {
	return 0, fmt.Errorf("erreur inscription utilisateur: %w", err)
	}

	retrun result.LastInsertId()
}