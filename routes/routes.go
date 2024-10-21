package routes

import (
	"expensetracker/controllers"
	"expensetracker/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// User routes
	api.Post("/users", controllers.CreateUser)
	api.Post("/login", controllers.Login)

	// Protected routes
	api.Use(middlewares.Protected())

	api.Get("/users/:id", controllers.GetUser)
	api.Post("/expenses", controllers.AddExpense)
	api.Get("/expenses/:id", controllers.GetUserExpenses)
	api.Get("/expenses", controllers.GetAllExpenses)
	api.Get("/download-balance-sheet", controllers.DownloadBalanceSheet)
}
