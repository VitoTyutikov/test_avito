package service

import (
	"avito_test_task/models"
	"avito_test_task/repository"
)

type FeatureService struct {
	repo *repository.FeatureRepository
}

func NewFeaturesService() *FeatureService {
	return &FeatureService{
		repo: repository.NewFeatureRepository(),
	}
}

func (s *FeatureService) FindAll() ([]models.Feature, error) {
	return s.repo.FindAll()
}

func (s *FeatureService) FindByID(featureID uint64) (*models.Feature, error) {
	return s.repo.FindByID(featureID)
}

func (s *FeatureService) Create(featureRequest *models.FeatureRequestBody) error {
	return s.repo.Create(featureRequest)
}

func (s *FeatureService) Delete(featureID uint64) error {
	return s.repo.Delete(featureID)
}
