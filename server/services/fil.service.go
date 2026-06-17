package services

import (
	"fmt"
	"projet_forum/models"
	"projet_forum/repositories"
)

type FilService struct {
	filRepository *repositories.FilRepository
}

func InitFilService(FilRepository *repositories.FilRepository) *FilService {
	return &FilService{filRepository: FilRepository}
}

func (s *FilService) Create(fil models.Fil) (int, error) {
	if fil.Titre == "" || fil.User_c.Id == 0 || fil.Tag_c.Id == 0 {
		return -1, fmt.Errorf(" Erreur ajout fil - Données manquantes ou invalides")
	}

	filId, filErr := s.filRepository.CreateFil(fil)
	if filErr != nil {
		return -1, filErr
	}

	return filId, nil
}

func (s *FilService) CreateMessage(message models.Message) (int, error) {
	if message.Contenu == "" || message.User_c.Id == 0 || message.Fil_c.Id == 0 {
		return -1, fmt.Errorf(" Erreur ajout message - Données manquantes ou invalides")
	}

	messageId, messageErr := s.filRepository.CreateMessageFil(message)
	if messageErr != nil {
		return -1, messageErr
	}

	return messageId, nil
}

func (s *FilService) ReadAll() ([]models.Fil, error) {
	filsList, filsErr := s.filRepository.ReadAll()
	if filsErr != nil {
		return nil, filsErr
	}

	return filsList, nil
}

func (s *FilService) ReadById(idFil int) (models.Fil, error) {
	if idFil <= 0 {
		return models.Fil{}, fmt.Errorf(" Erreur récupération fil - identifiant invalide : %d", idFil)
	}

	fil, filErr := s.filRepository.ReadById(idFil)
	if filErr != nil {
		return models.Fil{}, filErr
	}

	return fil, nil
}

func (s *FilService) FilByIdMessages(idFil int, tri string) ([]models.Message, error) {
	if idFil <= 0 {
		return []models.Message{}, fmt.Errorf(" Erreur récupération produit - identifiant invalide : %d", idFil)
	}

	fil, filErr := s.filRepository.FilByIdMessages(idFil, tri)
	if filErr != nil {
		return []models.Message{}, filErr
	}

	return fil, nil
}

func (s *FilService) FilsPetanque() ([]models.Fil, error) {
	filsList, filsErr := s.filRepository.FilsPetanque()
	if filsErr != nil {
		return nil, filsErr
	}

	return filsList, nil
}

func (s *FilService) FilsCuisine() ([]models.Fil, error) {
	filsList, filsErr := s.filRepository.FilsCuisine()
	if filsErr != nil {
		return nil, filsErr
	}

	return filsList, nil
}

func (s *FilService) FilsNature() ([]models.Fil, error) {
	filsList, filsErr := s.filRepository.FilsNature()
	if filsErr != nil {
		return nil, filsErr
	}

	return filsList, nil
}

func (s *FilService) AjoutReaction(reaction models.Reaction) (int, error) {
	if reaction.User_c.Id == 0 || reaction.Message_c.Id == 0 || reaction.Type_reac == "" {
		return -1, fmt.Errorf(" Erreur réaction - Données manquantes ou invalides")
	}

	if reaction.Type_reac != "like" && reaction.Type_reac != "dislike" {
		return -1, fmt.Errorf(" Erreur réaction - Type de réaction invalide ('like' ou 'dislike')")
	}

	reac, reacErr := s.filRepository.AjoutReaction(reaction)
	if reacErr != nil {
		return -1, reacErr
	}
	return reac, nil
}

func (s *FilService) GetFilOwner(idFil int) (int, error) {
	return s.filRepository.GetFilOwner(idFil)
}

func (s *FilService) UpdateFilById(fil models.Fil) error {
	if fil.Id == 0 || fil.Titre == "" || fil.Tag_c.Id == 0 {
		return fmt.Errorf(" Erreur modification fil - Donnees manquantes ou invalides")
	}

	return s.filRepository.UpdateFilById(fil)
}
