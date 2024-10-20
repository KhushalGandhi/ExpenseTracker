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
	if err := services.AddExpense(expense); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add expense"})
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
