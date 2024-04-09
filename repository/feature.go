package repository

import (
	"avito_test_task/db"
	"avito_test_task/models"
	"gorm.io/gorm"
)

type FeatureRepository struct {
	DB *gorm.DB
}

func NewFeatureRepository() *FeatureRepository {
	return &FeatureRepository{
		DB: db.DB,
	}
}

func (r *FeatureRepository) FindAll() ([]models.Feature, error) {
	var features []models.Feature
	err := r.DB.Find(&features).Error
	return features, err
}

func (r *FeatureRepository) FindByID(featureID uint) (*models.Feature, error) {
	var feature models.Feature
	err := r.DB.Where("feature_id = ?", featureID).First(&feature).Error
	return &feature, err
}

func (r *FeatureRepository) Create(featureRequest *models.FeatureRequestBody) error {
	feature := models.Feature{
		Description: featureRequest.Description,
	}
	return r.DB.Create(feature).Error
}

func (r *FeatureRepository) Delete(featureID uint) error {
	return r.DB.Delete(&models.FeatureRequestBody{}, featureID).Error
}
