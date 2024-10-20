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
}

type ExpenseSplit struct {
	gorm.Model
	ExpenseID  uint    `json:"expense_id"`
	UserID     uint    `json:"user_id"`
	Amount     float64 `json:"amount"`
	Percentage float64 `json:"percentage"`
}
