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

func (s *FilService) ReadAll() ([]models.Fil, error) {
	filsList, filsErr := s.filRepository.ReadAll()
	if filsErr != nil {
		return nil, filsErr
	}

	return filsList, nil
}

func (s *FilService) FilByIdMessages(idFil int) ([]models.Message, error) {
	if idFil <= 0 {
		return []models.Message{}, fmt.Errorf(" Erreur récupération produit - identifiant invalide : %d", idFil)
	}

	fil, filErr := s.filRepository.FilByIdMessages(idFil)
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
