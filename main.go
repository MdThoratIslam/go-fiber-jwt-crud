package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-jwt-crud/database"
	"go-fiber-jwt-crud/routes"
)

func main() {
	app := fiber.New()

	// Connect to Database
	database.ConnectDB()

	// Setup Routes
	routes.SetupRoutes(app)

	// Start Server
	app.Listen(":8080")
}
