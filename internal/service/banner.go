package service

import (
	"avito_test_task/internal/models"
	"avito_test_task/internal/repository"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

type BannerService struct {
	bannerRepo *repository.BannerRepository
}

func NewBannerService() *BannerService {
	return &BannerService{
		bannerRepo: repository.NewBannerRepository(),
	}
}

func (s *BannerService) IsBannerWithFeatureAndTagExists(featureID uint64, tagIDs []uint64, bannerID uint64) (bool, error) {
	return s.bannerRepo.IsBannerWithFeatureAndTagExistsWithoutBannerID(featureID, tagIDs, bannerID)
}

func (s *BannerService) FindByID(id uint64) (*models.Banner, error) {
	return s.bannerRepo.FindByID(id)
}

func (s *BannerService) UpdateBannerWithTags(oldBanner *models.Banner, newBanner *models.BannerRequestBody) (int, error) {
	return s.bannerRepo.UpdateBannerWithTags(oldBanner, newBanner)
}

func (s *BannerService) DeleteByID(id uint64) *gorm.DB {
	return s.bannerRepo.DeleteByID(id)
}

func (s *BannerService) GetBanners(featureID uint64, tagID uint64, limit int, offset int) ([]models.BannerResponseBody, error) {
	return s.bannerRepo.GetBanners(featureID, tagID, limit, offset)
}

func (s *BannerService) GetUserBanner(featureId uint64, tagId uint64, token string) (models.Banner, error) {
	return s.bannerRepo.GetUserBanner(featureId, tagId, token)
}

func (s *BannerService) CreateBannerWithTags(request *models.BannerRequestBody) (models.Banner, error) {
	return s.bannerRepo.CreateBannerWithTags(request)
}

func (s *BannerService) Update(id uint64, request *models.BannerRequestBody) (int, error) {
	oldBanner, err := s.FindByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusNotFound, nil
	} else if err != nil {
		return http.StatusInternalServerError, err
	}
	// check no new tags which exists with feature_id added to banner
	bannerExists, err := s.IsBannerWithFeatureAndTagExists(request.FeatureID, request.TagIds, oldBanner.BannerID)
	if bannerExists {
		return http.StatusBadRequest, errors.New("banner with this tag_id and feature_id already exists")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return s.UpdateBannerWithTags(oldBanner, request)

}
