package service

import (
	"avito_test_task/models"
	"avito_test_task/repository"
)

type TagService struct {
	repo *repository.TagRepository
}

func NewTagService() *TagService {
	return &TagService{
		repo: repository.NewTagRepository(),
	}
}

func (s *TagService) FindAll() ([]models.Tag, error) {
	return s.repo.FindAll()
}

func (s *TagService) FindByID(tagID uint) (*models.Tag, error) {
	return s.repo.FindByID(tagID)
}

func (s *TagService) Create(tagRequest *models.TagRequestBody) error {
	return s.repo.Create(tagRequest)
}

func (s *TagService) Delete(tagID uint) error {
	return s.repo.Delete(tagID)
}
