package models

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model

	Name string `gorm:"type:varchar(50);uniqueIndex"`
}

type VehicleResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
