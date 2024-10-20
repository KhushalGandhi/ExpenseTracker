package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email  string `gorm:"unique;not null" json:"email"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
}
