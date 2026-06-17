package services

import (
	"client/api"
	"client/dto"
	"fmt"
)

type FilService struct {
	filApi *api.FilApi
}

func InitFilService(filApi *api.FilApi) *FilService {
	return &FilService{filApi: filApi}
}

func (s *FilService) Create(fil dto.FilDto, token string) (int, error) {
	if fil.Titre == "" || fil.Tag_c.Id == 0 {
		return -1, fmt.Errorf(" Erreur ajout fil - Données manquantes ou invalides")
	}

	filId, _, err := s.filApi.Create(fil, token)
	if err != nil {
		return -1, err
	}

	return filId, nil
}

func (s *FilService) CreateMessage(message dto.MessageDto, idFil int, token string) (int, error) {
	if message.Contenu == "" {
		return -1, fmt.Errorf(" Erreur ajout message - Données manquantes ou invalides")
	}

	messageId, _, err := s.filApi.CreateMessage(message, idFil, token)
	if err != nil {
		return -1, err
	}

	return messageId, nil
}

func (s *FilService) ReadAll() ([]dto.FilDto, error) {
	return s.filApi.ReadAll()
}

func (s *FilService) ReadById(idFil int) (dto.FilDto, error) {
	if idFil <= 0 {
		return dto.FilDto{}, fmt.Errorf("Erreur récupération fil - identifiant invalide : %d", idFil)
	}

	return s.filApi.ReadById(idFil)
}

func (s *FilService) ReadByIdMessages(idFil int, tri string) ([]dto.MessageDto, error) {
	if idFil <= 0 {
		return []dto.MessageDto{}, fmt.Errorf("Erreur récupération produit - identifiant invalide : %d", idFil)
	}

	return s.filApi.ReadByIdMessages(idFil, tri)
}

func (s *FilService) FilsPetanque() ([]dto.FilDto, error) {
	return s.filApi.FilsPetanque()
}

func (s *FilService) FilsCuisine() ([]dto.FilDto, error) {
	return s.filApi.FilsCuisine()
}

func (s *FilService) FilsNature() ([]dto.FilDto, error) {
	return s.filApi.FilsNature()
}

func (s *FilService) AjoutReaction(reaction dto.Reaction, token string) (int, error) {
	if reaction.Message_c.Id == 0 || reaction.Type_reac == "" {
		return -1, fmt.Errorf(" Erreur reaction - Données manquantes ou invalides")
	}

	if reaction.Type_reac != "like" && reaction.Type_reac != "dislike" {
		return -1, fmt.Errorf(" Erreur réaction - Type de réaction invalide")
	}
	reac, _, err := s.filApi.AjoutReaction(reaction, token)
	if err != nil {
		return -1, err
	}

	return reac, nil
}

func (s *FilService) UpdateFilById(fil dto.FilDto, token string) error {
	if fil.Id <= 0 {
		return fmt.Errorf("Erreur modification fil - Identifiant invalide")
	}
	if fil.Titre == "" || fil.Tag_c.Id == 0 {
		return fmt.Errorf("Erreur modification fil - Données manquantes ou invalides")
	}

	return s.filApi.UpdateFilById(fil, token)
}
