package services

import (
	"client/api"
	"client/dto"
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
