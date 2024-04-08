package repository

import (
	"avito_test_task/db"
	"avito_test_task/models"
	"gorm.io/gorm"
)

type TagsRepository struct {
	DB *gorm.DB
}

func NewTagsRepository() *TagsRepository {
	return &TagsRepository{
		DB: db.DB,
	}
}

func (r *TagsRepository) FindAll() ([]models.Tags, error) {
	var tags []models.Tags
	err := r.DB.Find(&tags).Error
	return tags, err
}

func (r *TagsRepository) FindByID(tagID uint) (*models.Tags, error) {
	var tag models.Tags
	err := r.DB.Where("tag_id = ?", tagID).First(&tag).Error
	return &tag, err
}

func (r *TagsRepository) Create(tagRequest *models.TagsRequestBody) error {
	tag := models.Tags{
		Description: tagRequest.Description,
	}
	return r.DB.Create(tag).Error
}

func (r *TagsRepository) Delete(tagID uint) error {
	return r.DB.Delete(&models.Tags{}, tagID).Error
}
