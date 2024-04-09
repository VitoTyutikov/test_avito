package models

type Feature struct {
	FeatureID   uint   `json:"feature_id" gorm:"primarykey"`
	Description string `json:"description"`
}

type FeatureRequestBody struct {
	Description string `json:"description"`
}
