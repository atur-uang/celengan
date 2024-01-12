package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	UserName string `gorm:"type:varchar(50);uniqueIndex"`
	FullName *string
	Email    *string
	Phone    string `gorm:"index"`
	Password string
}
