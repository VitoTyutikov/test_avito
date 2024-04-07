package models

import (
	"encoding/json"
	"time"
)

type Banner struct {
	BannerID  uint            `json:"banner_id" gorm:"primarykey"`
	FeatureID int             `json:"feature_id"`
	Content   json.RawMessage `json:"content" gorm:"type:json"` // For JSON type in postgres, RawMessage is []byte
	IsActive  bool            `json:"is_active"`
	CreatedAt time.Time       `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"type:timestamp"`
}

type BannerRequestBody struct {
	TagIds    []int           `json:"tag_ids"`
	FeatureID int             `json:"feature_id"`
	Content   json.RawMessage `json:"content"` // For JSON type in postgres, RawMessage is []byte
	IsActive  bool            `json:"is_active"`
}
