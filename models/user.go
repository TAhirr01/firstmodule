package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Age      uint32
	Email    string
	Password string
}
