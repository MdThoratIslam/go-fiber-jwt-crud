package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go-fiber-jwt-crud/database"
	logger "go-fiber-jwt-crud/log"
	"go-fiber-jwt-crud/routes"
	"os"
)

func main() {

	app := fiber.New()

	if err := godotenv.Load(); err != nil {
		logger.Error("Error loading .env file", err)
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// server run successfully messagestore in log file
	logger.Success("Server is running on port: " + os.Getenv("APP_PORT"))
	// Connect to Database
	database.ConnectDB()
	// Enable CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("FRONTEND_URL"), // Use FRONTEND_URL from .env
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	// Setup Routes
	routes.SetupRoutes(app)
	// app_host from .env
	// app_port from .env
	// db_name from .env

	// Start Server
	//app.Listen(":3000")
	env := godotenv.Load()
	if env != nil {
		logger.Error("Error loading .env file", env)
	}
	app_host := os.Getenv("APP_HOST")
	app_port := os.Getenv("APP_PORT")
	app.Listen(app_host + ":" + app_port)
}
