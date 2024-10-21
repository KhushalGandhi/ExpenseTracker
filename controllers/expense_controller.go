package controllers

import (
	"expensetracker/models"
	"expensetracker/services"
	"github.com/gofiber/fiber/v2"
)

func AddExpense(c *fiber.Ctx) error {
	expense := new(models.Expense)

	if err := c.BodyParser(expense); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Get the user_id from the token
	userIdFloat, ok := c.Locals("user_id").(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	// Convert userIdFloat to uint
	userId := uint(userIdFloat)

	// Use userId in your expense creation logic
	expense.UserID = userId // Assuming your Expense model has a UserID field

	if err := services.AddExpense(expense); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create expense"})
	}
	return c.Status(fiber.StatusCreated).JSON(expense)
}

func GetUserExpenses(c *fiber.Ctx) error {
	id := c.Params("id")
	expenses, err := services.GetUserExpenses(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch expenses"})
	}
	return c.JSON(expenses)
}

func GetAllExpenses(c *fiber.Ctx) error {
	expenses, err := services.GetAllExpenses()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch expenses"})
	}
	return c.JSON(expenses)
}

func DownloadBalanceSheet(c *fiber.Ctx) error {
	balanceSheet, err := services.GenerateBalanceSheetCSV()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate balance sheet"})
	}

	// Send the CSV as a file download
	c.Set("Content-Disposition", "attachment; filename=balance_sheet.csv")
	c.Set("Content-Type", "text/csv")
	return c.SendStream(balanceSheet)
}
