package services

import "projet_forum/repositories"

type FilService struct {
	filRepository *repositories.FilRepository
}

func InitFilService(FilRepository *repositories.FilRepository) *FilService {
	return &FilService{filRepository: FilRepository}
}
