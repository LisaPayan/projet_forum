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

	sqlResult, sqlErr := r.db.Query("SELECT f.id, f.titre, f.statut, f.score, u.pseudo FROM fils f INNER JOIN users u ON f.fk_user = u.id WHERE f.statut != 2;")
	if sqlErr != nil {
		return listFils, fmt.Errorf("Erreur récupération fil - Erreur: \n\t %s", sqlErr.Error())
	}
	defer sqlResult.Close()
	for sqlResult.Next() {
		var fil models.Fil

		errScan := sqlResult.Scan(&fil.Id, &fil.Titre, &fil.Statut, &fil.Score, &fil.User.Pseudo)
		if errScan != nil {
			continue
		}
		listFils = append(listFils, fil)
	}
	return listFils, nil
}

func (r *FilRepository) FilByIdMessages(idFil int) ([]models.Message, error) {
	var listMessages []models.Message

	query := "SELECT m.id, m.contenu, m.is_published, m.date_publication, m.nbr_dislikes, u.pseudo, f.Titre FROM messages m INNER JOIN users u ON m.fk_user = u.id INNER JOIN fils f ON m.fk_fil = f.id WHERE m.fk_fil = ? AND m.is_published = 1 ORDER BY m.date_publication ASC; "
	// m.nbr_likes,
	sqlResult, sqlErr := r.db.Query(query, idFil)
	if sqlErr != nil {
		return listMessages, fmt.Errorf("Erreur récupération messages - Erreur: \n\t %s", sqlErr.Error())
	}
	defer sqlResult.Close()
	for sqlResult.Next() {
		var message models.Message

		errScan := sqlResult.Scan(&message.Id, &message.Contenu, &message.IsPublished, &message.PublishedAt, &message.NbrDislikes, &message.User_c.Pseudo, &message.Fil.Titre)
		// &message.NbrLikes,
		if errScan != nil {
			continue
		}
		listMessages = append(listMessages, message)
	}
	return listMessages, nil
}
