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

func (s *FilService) ReadAll() ([]dto.FilDto, error) {
	return s.filApi.ReadAll()
}

func (s *FilService) ReadByIdMessages(idFil int) ([]dto.MessageDto, error) {
	if idFil <= 0 {
		return []dto.MessageDto{}, fmt.Errorf("Erreur récupération produit - identifiant invalide : %d", idFil)
	}

	return s.filApi.ReadByIdMessages(idFil)
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
