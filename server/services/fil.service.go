package services

import (
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
