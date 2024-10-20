package migrations

import (
	"expensetracker/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Expense{}, &models.ExpenseSplit{})
}
