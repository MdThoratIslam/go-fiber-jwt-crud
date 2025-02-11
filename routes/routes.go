package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-jwt-crud/controllers"
	"go-fiber-jwt-crud/middleware"
)

func SetupRoutes(app *fiber.App) {
	// Authentication Routes
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)

	// User CRUD Routes (JWT Protected)
	api := app.Group("/api", middleware.JWTMiddleware())
	api.Get("/users", controllers.GetUsers)
	api.Get("/users/:id", controllers.GetUser)
	api.Put("/users/:id", controllers.UpdateUser)
	api.Delete("/users/:id", controllers.DeleteUser)

	api.Get("/logs", controllers.GetLog)
	api.Post("/logout", middleware.Logout)
}
