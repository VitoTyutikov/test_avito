package models

type Features struct {
	FeatureID   uint   `json:"feature_id" gorm:"primarykey"`
	Description string `json:"description"`
}
