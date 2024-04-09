package models

import (
	"encoding/json"
	"time"
)

type Banner struct {
	BannerID  uint            `json:"banner_id" gorm:"primarykey"`
	FeatureID uint            `json:"feature_id"`
	Content   json.RawMessage `json:"content" gorm:"type:json"` // For JSON type in postgres, RawMessage is []byte
	IsActive  bool            `json:"is_active"`
	CreatedAt time.Time       `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"type:timestamp"`
	Feature   Feature         `gorm:"references:FeatureID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type BannerRequestBody struct {
	TagIds    []uint          `json:"tag_ids"`
	FeatureID uint            `json:"feature_id"`
	Content   json.RawMessage `json:"content"` // For JSON type in postgres, RawMessage is []byte
	IsActive  bool            `json:"is_active"`
}

type BannerCreatePayload struct {
	FeatureID uint            `json:"feature_id"`
	Content   json.RawMessage `json:"content"` // For JSON type in postgres, RawMessage is []byte
	IsActive  bool            `json:"is_active"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
