package models

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model

	Name string `gorm:"uniqueIndex"`
}

type VehicleAPI struct {
	ID   uint
	Name string
}
