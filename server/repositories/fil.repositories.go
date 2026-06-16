package repositories

import (
	"database/sql"
	"fmt"
	"projet_forum/models"
	"time"
)

type FilRepository struct {
	db *sql.DB
}

func InitFilRepository(db *sql.DB) *FilRepository {
	return &FilRepository{db}
}

func (r *FilRepository) CreateFil(fil models.Fil) (int, error) {
	query := "INSERT INTO `fils` (`titre`, `fk_user`, `fk_tag`) VALUES (?,?,?);"

	sqlResult, sqlErr := r.db.Exec(query,
		fil.Titre,
		fil.User_c.Id,
		fil.Tag_c.Id,
	)

	if sqlErr != nil {
		return -1, fmt.Errorf(" Erreur ajout fil - Erreur : \n\t %s", sqlErr.Error())
	}

	id, idErr := sqlResult.LastInsertId()
	if idErr != nil {
		return -1, fmt.Errorf(" Erreur ajout fil - Erreur récupération identifiant : \n\t %s", idErr.Error())
	}

	return int(id), nil
}

func (r *FilRepository) ReadAll() ([]models.Fil, error) {
	var listFils []models.Fil

	sqlResult, sqlErr := r.db.Query("SELECT f.id, f.titre, f.statut, f.score, u.pseudo, t.id, t.nom FROM fils f LEFT JOIN users u ON f.fk_user = u.id LEFT JOIN tags t ON f.fk_tag = t.id WHERE f.statut != 'archivé';")

	if sqlErr != nil {
		return listFils, fmt.Errorf("Erreur récupération fil - Erreur: \n\t %s", sqlErr.Error())
	}
	defer sqlResult.Close()
	for sqlResult.Next() {
		var fil models.Fil

		errScan := sqlResult.Scan(&fil.Id, &fil.Titre, &fil.Statut, &fil.Score, &fil.User_c.Pseudo, &fil.Tag_c.Id, &fil.Tag_c.Nom)
		if errScan != nil {
			continue
		}
		listFils = append(listFils, fil)
	}
	return listFils, nil
}

func (r *FilRepository) FilByIdMessages(idFil int) ([]models.Message, error) {
	var listMessages []models.Message

	query := "SELECT m.id, m.contenu, m.date_publication, u.pseudo, f.titre, t.id FROM messages m LEFT JOIN users u ON m.fk_user = u.id LEFT JOIN fils f ON m.fk_fil = f.id LEFT JOIN tags t ON f.fk_tag = t.id WHERE m.fk_fil = ? AND f.statut != 'archivé' ORDER BY m.date_publication ASC; "

	sqlResult, sqlErr := r.db.Query(query, idFil)
	if sqlErr != nil {
		return listMessages, fmt.Errorf("Erreur récupération messages - Erreur: \n\t %s", sqlErr.Error())
	}
	defer sqlResult.Close()
	for sqlResult.Next() {
		var message models.Message

		errScan := sqlResult.Scan(&message.Id, &message.Contenu, &message.PublishedAt, &message.User_c.Pseudo, &message.Fil_c.Titre, &message.Fil_c.Tag_c.Id)
		if errScan != nil {
			continue
		}
		listMessages = append(listMessages, message)
	}
	return listMessages, nil
}

func (r *FilRepository) FilByIdMessagesAnciens(idFil int) ([]models.Message, error) {
	var listMessages []models.Message

	query := "SELECT m.id, m.contenu, m.date_publication, u.pseudo, f.titre, t.id FROM messages m LEFT JOIN users u ON m.fk_user = u.id LEFT JOIN fils f ON m.fk_fil = f.id LEFT JOIN tags t ON f.fk_tag = t.id WHERE m.fk_fil = ? AND f.statut != 'archivé' ORDER BY m.date_publication ASC; "

	sqlResult, sqlErr := r.db.Query(query, idFil)
	if sqlErr != nil {
		return listMessages, fmt.Errorf("Erreur récupération messages - Erreur: \n\t %s", sqlErr.Error())
	}
	defer sqlResult.Close()
	for sqlResult.Next() {
		var message models.Message

		errScan := sqlResult.Scan(&message.Id, &message.Contenu, &message.PublishedAt, &message.User_c.Pseudo, &message.Fil_c.Titre, &message.Fil_c.Tag_c.Id)
		if errScan != nil {
			continue
		}
		listMessages = append(listMessages, message)
	}
	return listMessages, nil
}

func (r *FilRepository) FilByIdMessagesRecents(idFil int) ([]models.Message, error) {
	var listMessages []models.Message

	query := "SELECT m.id, m.contenu, m.date_publication, u.pseudo, f.titre, t.id FROM messages m LEFT JOIN users u ON m.fk_user = u.id LEFT JOIN fils f ON m.fk_fil = f.id LEFT JOIN tags t ON f.fk_tag = t.id WHERE m.fk_fil = ? AND f.statut != 'archivé' ORDER BY m.date_publication DESC; "

	sqlResult, sqlErr := r.db.Query(query, idFil)
	if sqlErr != nil {
		return listMessages, fmt.Errorf("Erreur récupération messages - Erreur: \n\t %s", sqlErr.Error())
	}
	defer sqlResult.Close()
	for sqlResult.Next() {
		var message models.Message

		errScan := sqlResult.Scan(&message.Id, &message.Contenu, &message.PublishedAt, &message.User_c.Pseudo, &message.Fil_c.Titre, &message.Fil_c.Tag_c.Id)
		if errScan != nil {
			continue
		}
		listMessages = append(listMessages, message)
	}
	return listMessages, nil
}

func (r *FilRepository) CreateMessageFil(message models.Message) (int, error) {
	query := "INSERT INTO `messages` (`contenu`, `date_publication`, `fk_user`, `fk_fil`) VALUES (?,?,?,?);"

	sqlResult, sqlErr := r.db.Exec(query,
		message.Contenu,
		time.Now().Format("2006-01-02 15:04:05"),
		message.User_c.Id,
		message.Fil_c.Id,
	)

	if sqlErr != nil {
		return -1, fmt.Errorf(" Erreur ajout message - Erreur : \n\t %s", sqlErr.Error())
	}

	id, idErr := sqlResult.LastInsertId()
	if idErr != nil {
		return -1, fmt.Errorf(" Erreur ajout message - Erreur récupération identifiant : \n\t %s", idErr.Error())
	}

	return int(id), nil
}

func (r *FilRepository) FilsPetanque() ([]models.Fil, error) {
	var listFils []models.Fil

	sqlResult, sqlErr := r.db.Query("SELECT f.id, f.titre, f.statut, f.score, u.pseudo, t.id, t.nom, t.description FROM fils f LEFT JOIN users u ON f.fk_user = u.id LEFT JOIN tags t ON f.fk_tag = t.id WHERE f.statut != 'archivé' AND t.id = 1;")

	if sqlErr != nil {
		return listFils, fmt.Errorf("Erreur récupération fil - Erreur: \n\t %s", sqlErr.Error())
	}
	defer sqlResult.Close()
	for sqlResult.Next() {
		var fil models.Fil

		errScan := sqlResult.Scan(&fil.Id, &fil.Titre, &fil.Statut, &fil.Score, &fil.User_c.Pseudo, &fil.Tag_c.Id, &fil.Tag_c.Nom, &fil.Tag_c.Description)
		if errScan != nil {
			continue
		}
		listFils = append(listFils, fil)
	}
	return listFils, nil
}

func (r *FilRepository) FilsCuisine() ([]models.Fil, error) {
	var listFils []models.Fil

	sqlResult, sqlErr := r.db.Query("SELECT f.id, f.titre, f.statut, f.score, u.pseudo, t.id, t.nom, t.description FROM fils f LEFT JOIN users u ON f.fk_user = u.id LEFT JOIN tags t ON f.fk_tag = t.id WHERE f.statut != 'archivé' AND t.id = 2;")

	if sqlErr != nil {
		return listFils, fmt.Errorf("Erreur récupération fil - Erreur: \n\t %s", sqlErr.Error())
	}
	defer sqlResult.Close()
	for sqlResult.Next() {
		var fil models.Fil

		errScan := sqlResult.Scan(&fil.Id, &fil.Titre, &fil.Statut, &fil.Score, &fil.User_c.Pseudo, &fil.Tag_c.Id, &fil.Tag_c.Nom, &fil.Tag_c.Description)
		if errScan != nil {
			continue
		}
		listFils = append(listFils, fil)
	}
	return listFils, nil
}

func (r *FilRepository) FilsNature() ([]models.Fil, error) {
	var listFils []models.Fil

	sqlResult, sqlErr := r.db.Query("SELECT f.id, f.titre, f.statut, f.score, u.pseudo, t.id, t.nom, t.description FROM fils f LEFT JOIN users u ON f.fk_user = u.id LEFT JOIN tags t ON f.fk_tag = t.id WHERE f.statut != 'archivé' AND t.id = 3;")

	if sqlErr != nil {
		return listFils, fmt.Errorf("Erreur récupération fil - Erreur: \n\t %s", sqlErr.Error())
	}
	defer sqlResult.Close()
	for sqlResult.Next() {
		var fil models.Fil

		errScan := sqlResult.Scan(&fil.Id, &fil.Titre, &fil.Statut, &fil.Score, &fil.User_c.Pseudo, &fil.Tag_c.Id, &fil.Tag_c.Nom, &fil.Tag_c.Description)
		if errScan != nil {
			continue
		}
		listFils = append(listFils, fil)
	}
	return listFils, nil
}

// func (r *FilRepository) Like(idMess int) ([]models.Reaction, error) {
// 	var listMessages []models.Message

// 	query := "SELECT m.id, m.contenu, m.date_publication, u.pseudo, f.titre, t.id FROM messages m LEFT JOIN users u ON m.fk_user = u.id LEFT JOIN fils f ON m.fk_fil = f.id LEFT JOIN tags t ON f.fk_tag = t.id WHERE m.fk_fil = ? AND f.statut != 'archivé' ORDER BY m.date_publication ASC; "

// 	sqlResult, sqlErr := r.db.Query(query, idMess)
// 	if sqlErr != nil {
// 		return listMessages, fmt.Errorf("Erreur récupération messages - Erreur: \n\t %s", sqlErr.Error())
// 	}
// 	defer sqlResult.Close()
// 	for sqlResult.Next() {
// 		var message models.Message

// 		errScan := sqlResult.Scan(&message.Id, &message.Contenu, &message.PublishedAt, &message.User_c.Pseudo, &message.Fil_c.Titre, &message.Fil_c.Tag_c.Id)
// 		if errScan != nil {
// 			continue
// 		}
// 		listMessages = append(listMessages, message)
// 	}
// 	return listMessages, nil
// }
