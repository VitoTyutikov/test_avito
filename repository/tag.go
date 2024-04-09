package repository

import (
	"avito_test_task/db"
	"avito_test_task/models"
	"gorm.io/gorm"
)

type TagRepository struct {
	DB *gorm.DB
}

func NewTagRepository() *TagRepository {
	return &TagRepository{
		DB: db.DB,
	}
}

func (r *TagRepository) FindAll() ([]models.Tag, error) {
	var tags []models.Tag
	err := r.DB.Find(&tags).Error
	return tags, err
}

func (r *TagRepository) FindByID(tagID uint64) (*models.Tag, error) {
	var tag models.Tag
	err := r.DB.Where("tag_id = ?", tagID).First(&tag).Error
	return &tag, err
}

func (r *TagRepository) Create(tagRequest *models.TagRequestBody) error {
	tag := models.Tag{
		Description: tagRequest.Description,
	}
	return r.DB.Create(tag).Error
}

func (r *TagRepository) Delete(tagID uint64) error {
	return r.DB.Delete(&models.Tag{}, tagID).Error
}
