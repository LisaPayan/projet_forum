package repositories

import (
	"database/sql"
	"fmt"
	"projet_forum/models"
)

type FilRepository struct {
	db *sql.DB
}

func InitFilRepository(db *sql.DB) *FilRepository {
	return &FilRepository{db}
}

func (r *FilRepository) ReadAll() ([]models.Fil, error) {
	var listFils []models.Fil

	sqlResult, sqlErr := r.db.Query("SELECT id, titre FROM fils;")
	if sqlErr != nil {
		return listFils, fmt.Errorf("Erreur récupération fil - Erreur: \n\t %s", sqlErr.Error())
	}
	defer sqlResult.Close()
	for sqlResult.Next() {
		var fil models.Fil

		errScan := sqlResult.Scan(&fil.Id, &fil.Titre)
		if errScan != nil {
			continue
		}
		listFils = append(listFils, fil)
	}
	return listFils, nil
}
