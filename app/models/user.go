package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	UserName string `gorm:"uniqueIndex"`
	FullName *string
	Email    *string
	Phone    string `gorm:"index"`
	Password string
}
