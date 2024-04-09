package service

import (
	"avito_test_task/models"
	"avito_test_task/repository"
)

type BannerService struct {
	repo *repository.BannerRepository
}

func NewBannerService() *BannerService {
	return &BannerService{
		repo: repository.NewBannerRepository(),
	}
}

func (s *BannerService) CreateBanner(bannerReq *models.BannerRequestBody) error {
	// Business logic can be added here before saving the banner
	return s.repo.Create(bannerReq)
}

func (s *BannerService) FindByID(id uint) (*models.Banner, error) {
	// Additional business logic can be added here
	return s.repo.FindByID(id)
}

//func (s *BannerService) UpdateBanner(banner *models.Banner) error {
//	// Business logic for updating a banner
//	return s.repo.Update(banner)
//}

func (s *BannerService) DeleteByID(id uint) error {
	// Additional logic before deletion, if necessary
	return s.repo.DeleteByID(id)
}

func (s *BannerService) FindAll() ([]models.Banner, error) {
	// Logic to handle retrieval of all banners
	return s.repo.FindAll()
}
