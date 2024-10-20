package services

import (
	"errors"
	"expensetracker/models"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetDB(database *gorm.DB) {
	db = database
}

func CreateUser(user *models.User) error {
	return db.Create(user).Error
}

func GetUser(id string) (models.User, error) {
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return user, errors.New("User not found")
	}
	return user, nil
}
