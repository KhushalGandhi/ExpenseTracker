package main

import (
	"expensetracker/migrations"
	"expensetracker/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func initDB() {
	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatalf("Error loading env file: %v", err)
	}

	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") + " sslmode=disable"
	var errDb error
	db, errDb = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errDb != nil {
		log.Fatal("Failed to connect to the database")
	}
	migrations.RunMigrations(db)
}

func main() {
	app := fiber.New()

	// Initialize database
	initDB()

	// Setup routes
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
