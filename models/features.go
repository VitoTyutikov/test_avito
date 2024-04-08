package models

type Features struct {
	FeatureID   uint   `json:"feature_id" gorm:"primarykey"`
	Description string `json:"description"`
}

type FeaturesRequestBody struct {
	Description string `json:"description"`
}
