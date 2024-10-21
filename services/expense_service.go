package services

import (
	"bytes"
	"encoding/csv"
	"errors"
	"expensetracker/models"
	"gorm.io/gorm"
	"strconv"
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

func GetAllExpenses() ([]models.Expense, error) {
	var expenses []models.Expense
	if err := db.Find(&expenses).Error; err != nil {
		return expenses, err
	}
	return expenses, nil
}

func GenerateBalanceSheetCSV() (*bytes.Buffer, error) {
	var expenses []models.Expense
	err := db.Find(&expenses).Error
	if err != nil {
		return nil, err
	}

	// Create a new buffer to hold the CSV data
	b := new(bytes.Buffer)
	writer := csv.NewWriter(b)

	// Write CSV headers
	headers := []string{"Expense ID", "Expense Name", "Amount", "User", "Split Method"}
	if err := writer.Write(headers); err != nil {
		return nil, err
	}

	// Write each expense to the CSV
	for _, expense := range expenses {
		row := []string{
			strconv.FormatUint(uint64(expense.ID), 10), // Expense ID
			expense.Name, // Expense Name
			strconv.FormatFloat(expense.Amount, 'f', 2, 64), // Amount
			expense.Name,      // User Name
			expense.SplitType, // Split Method (Equal/Exact/Percentage)
		}
		if err := writer.Write(row); err != nil {
			return nil, err
		}
	}

	// Flush the writer to make sure all data is written to the buffer
	writer.Flush()

	// Check if any errors occurred while flushing
	if err := writer.Error(); err != nil {
		return nil, err
	}

	return b, nil
}
