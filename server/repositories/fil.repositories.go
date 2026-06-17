package repositories

import (
	"database/sql"
	"fmt"
	"projet_forum/models"
	"strings"
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

func (r *FilRepository) ReadById(id int) (models.Fil, error) {
	var fil models.Fil
	sqlErr := r.db.QueryRow("SELECT f.id, f.titre, f.statut, f.score, u.pseudo, t.id, t.nom FROM fils f LEFT JOIN users u ON f.fk_user = u.id LEFT JOIN tags t ON f.fk_tag = t.id WHERE f.statut != 'archivé' AND f.id = ?;", id).
		Scan(&fil.Id, &fil.Titre, &fil.Statut, &fil.Score, &fil.User_c.Pseudo, &fil.Tag_c.Id, &fil.Tag_c.Nom)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return models.Fil{}, nil
		}
		return models.Fil{}, fmt.Errorf(" Erreur récupération produit - Erreur : \n\t %s", sqlErr.Error())
	}

	return fil, nil
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

	query := `SELECT m.id,
	m.contenu,
	m.date_publication,
	u.pseudo,
	f.titre,
	t.id,
	COUNT(CASE WHEN r.type_reaction = 'like' THEN 1 END) AS nb_likes,
    COUNT(CASE WHEN r.type_reaction = 'dislike' THEN 1 END) AS nb_dislikes,
	(COUNT(CASE WHEN r.type_reaction = 'like' THEN 1 END) - COUNT(CASE WHEN r.type_reaction = 'dislike' THEN 1 END)) AS score_popularite
	FROM messages m 
	LEFT JOIN users u ON m.fk_user = u.id 
	LEFT JOIN fils f ON m.fk_fil = f.id 
	LEFT JOIN tags t ON f.fk_tag = t.id 
	LEFT JOIN reactions r ON m.id = r.fk_message
	WHERE m.fk_fil = ? AND f.statut != 'archivé' 
	GROUP BY m.id, m.contenu, m.date_publication, u.pseudo, f.titre, t.id
    ORDER BY m.date_publication DESC;
	`
	sqlResult, sqlErr := r.db.Query(query, idFil)
	if sqlErr != nil {
		return listMessages, fmt.Errorf("Erreur récupération messages - Erreur: \n\t %s", sqlErr.Error())
	}
	defer sqlResult.Close()
	for sqlResult.Next() {
		var message models.Message

		errScan := sqlResult.Scan(&message.Id, &message.Contenu, &message.PublishedAt, &message.User_c.Pseudo, &message.Fil_c.Titre, &message.Fil_c.Tag_c.Id, &message.NbLikes, &message.NbDislikes, &message.ScorePopularite)
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

func (r *FilRepository) AjoutReaction(reaction models.Reaction) (int, error) {
	var currentType string

	query := "SELECT type_reaction FROM reactions WHERE fk_user = ? AND fk_message = ?; "

	err := r.db.QueryRow(query,
		reaction.User_c.Id,
		reaction.Message_c.Id).Scan(&currentType)

	if err != nil {
		if err == sql.ErrNoRows {
			queryInsert := "INSERT INTO reactions (fk_user, fk_message, type_reaction) VALUES (?, ?, ?);"
			_, errInsert := r.db.Exec(queryInsert,
				reaction.User_c.Id,
				reaction.Message_c.Id,
				reaction.Type_reac)
			if errInsert != nil {
				return -1, fmt.Errorf("Erreur insert réaction: %v", errInsert)
			}
			return 1, nil
		}
		return -1, fmt.Errorf("Erreur base de données lors de la sélection: %v", err)
	}

	currentType = strings.TrimSpace(currentType)
	targetType := strings.TrimSpace(reaction.Type_reac)

	if currentType == targetType {
		queryDelete := "DELETE FROM reactions WHERE fk_user = ? AND fk_message = ?;"
		_, errDelete := r.db.Exec(queryDelete,
			reaction.User_c.Id,
			reaction.Message_c.Id)
		if errDelete != nil {
			return -1, fmt.Errorf("Erreur delete réaction: %v", errDelete)
		}
		return 0, nil
	}

	queryUpdate := "UPDATE reactions SET type_reaction = ? WHERE fk_user = ? AND fk_message = ?;"
	_, errUpdate := r.db.Exec(queryUpdate,
		targetType,
		reaction.User_c.Id,
		reaction.Message_c.Id)
	if errUpdate != nil {
		return -1, fmt.Errorf("Erreur update réaction: %v", errUpdate)
	}
	return 1, nil
}
