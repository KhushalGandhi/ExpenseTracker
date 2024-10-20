package routes

import (
	"expensetracker/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// User routes
	api.Post("/users", controllers.CreateUser)
	api.Get("/users/:id", controllers.GetUser)

	// Expense routes
	api.Post("/expenses", controllers.AddExpense)
	api.Get("/expenses/:id", controllers.GetUserExpenses)
}
