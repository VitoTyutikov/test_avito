package models

type Feature struct {
	FeatureID uint64 `json:"feature_id" gorm:"primarykey"`
}
