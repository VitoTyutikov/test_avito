package models

import (
	"encoding/json"
	"time"
)

type Banner struct {
	BannerID  uint64          `json:"banner_id" gorm:"primarykey"`
	FeatureID uint64          `json:"feature_id" gorm:"not null;"`
	Content   json.RawMessage `json:"content" gorm:"type:jsonb; not null"` // For JSON type in postgres, RawMessage is []byte
	IsActive  bool            `json:"is_active" gorm:"not null"`
	CreatedAt time.Time       `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"type:timestamp"`
	Feature   Feature         `gorm:"references:FeatureID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type BannerRequestBody struct {
	TagIds    []uint64        `json:"tag_ids,omitempty" binding:"required"`
	FeatureID uint64          `json:"feature_id,omitempty" binding:"required"`
	Content   json.RawMessage `json:"content,omitempty" binding:"required"`
	IsActive  *bool           `json:"is_active,omitempty" binding:"required"` // pointer because default value is false and when
	// try to update with false field it lead to error in ShouldBindJSON
}

type BannerResponseBody struct {
	BannerID  uint64          `json:"banner_id"`
	TagIds    []uint64        `json:"tag_ids"`
	FeatureID uint64          `json:"feature_id"`
	Content   json.RawMessage `json:"content"`
	IsActive  bool            `json:"is_active"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
