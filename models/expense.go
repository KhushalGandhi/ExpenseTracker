package models

import (
	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	SplitType   string  `json:"split_type"` // Equal, Exact, or Percentage
	UserID      uint    `json:"user_id"`
	Name        string  `json:"name"`
	User        User    `gorm:"foreignKey:UserID"` // Reference to the user who paid
	// Expense name

}

//type Expense struct {
//	gorm.Model
//	Name        string  `json:"name"`  // Expense name
//	Amount      float64 `json:"amount"` // Expense amount
//	UserID      uint    `json:"user_id"` // Reference to the user who paid
//	SplitMethod string  `json:"split_method"` // Split method: Equal, Exact, Percentage
//
//	User User `gorm:"foreignKey:UserID"` // Reference to the user who paid
//}

type ExpenseSplit struct {
	gorm.Model
	ExpenseID  uint    `json:"expense_id"`
	UserID     uint    `json:"user_id"`
	Amount     float64 `json:"amount"`
	Percentage float64 `json:"percentage"`
}
