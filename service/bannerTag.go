package service

import (
	"avito_test_task/models"
	"avito_test_task/repository"
	"gorm.io/gorm"
)

type BannerTagService struct {
	repo *repository.BannerTagRepository
}

func NewBannerTagService() *BannerTagService {
	return &BannerTagService{
		repo: repository.NewBannerTagRepository(),
	}
}

func (s *BannerTagService) Create(bannerTag *models.BannerTag) error {
	return s.repo.Create(bannerTag)
}

func (s *BannerTagService) FindAll() ([]models.BannerTag, error) {
	return s.repo.FindAll()
}

func (s *BannerTagService) FindByID(bannerID, tagID uint64) (*models.BannerTag, error) {
	return s.repo.FindByID(bannerID, tagID)
}
func (s *BannerTagService) DeleteByBannerID(bannerID uint64) *gorm.DB {
	return s.repo.DeleteByBannerID(bannerID)
}
