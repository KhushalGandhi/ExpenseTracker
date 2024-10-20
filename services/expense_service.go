package services

import (
	"errors"
	"expensetracker/models"
	"gorm.io/gorm"
)

func AddExpense(expense *models.Expense) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&expense).Error; err != nil {
			return err
		}

		var splits []models.ExpenseSplit
		switch expense.SplitType {
		case "Equal":
			splits = splitEqual(expense)
		case "Exact":
			splits = splitExact(expense)
		case "Percentage":
			splits = splitPercentage(expense)
		default:
			return errors.New("invalid split type")
		}

		for _, split := range splits {
			if err := tx.Create(&split).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func GetUserExpenses(id string) ([]models.Expense, error) {
	var expenses []models.Expense
	if err := db.Where("user_id = ?", id).Find(&expenses).Error; err != nil {
		return expenses, err
	}
	return expenses, nil
}

// Splitting logic
func splitEqual(expense *models.Expense) []models.ExpenseSplit {
	participants := getParticipants(expense)
	share := expense.Amount / float64(len(participants))

	var splits []models.ExpenseSplit
	for _, participant := range participants {
		splits = append(splits, models.ExpenseSplit{
			ExpenseID: expense.ID,
			UserID:    participant.ID,
			Amount:    share,
		})
	}
	return splits
}

func splitExact(expense *models.Expense) []models.ExpenseSplit {
	// Implement logic to split by exact amounts
	return []models.ExpenseSplit{}
}

func splitPercentage(expense *models.Expense) []models.ExpenseSplit {
	// Implement logic to split by percentage
	return []models.ExpenseSplit{}
}

func getParticipants(expense *models.Expense) []models.User {
	// Dummy function to get participants of the expense
	return []models.User{}
}
