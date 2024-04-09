package repository

import (
	"avito_test_task/db"
	"avito_test_task/models"
	"gorm.io/gorm"
)

type BannerTagRepository struct {
	DB *gorm.DB
}

func NewBannerTagRepository() *BannerTagRepository {
	return &BannerTagRepository{
		DB: db.DB,
	}
}

func (r *BannerTagRepository) Create(bannerTag *models.BannerTag) error {
	return r.DB.Create(bannerTag).Error
}

func (r *BannerTagRepository) FindByID(bannerID, tagID uint64) (*models.BannerTag, error) {
	var bannerTag models.BannerTag
	err := r.DB.Where("banner_id = ? AND tag_id = ?", bannerID, tagID).First(&bannerTag).Error
	return &bannerTag, err
}

func (r *BannerTagRepository) Delete(bannerID, tagID uint64) error {
	return r.DB.Where("banner_id = ? AND tag_id = ?", bannerID, tagID).Delete(&models.BannerTag{}).Error
}

func (r *BannerTagRepository) FindAll() ([]models.BannerTag, error) {
	var bannerTags []models.BannerTag
	err := r.DB.Find(&bannerTags).Error
	return bannerTags, err
}
