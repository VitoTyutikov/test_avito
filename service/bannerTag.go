package service

import (
	"avito_test_task/models"
	"avito_test_task/repository"
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

func (s *BannerTagService) Delete(bannerID, tagID uint64) error {
	return s.repo.Delete(bannerID, tagID)
}
