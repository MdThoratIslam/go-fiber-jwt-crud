package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go-fiber-jwt-crud/database"
	logger "go-fiber-jwt-crud/log"
	"go-fiber-jwt-crud/middleware"
	"go-fiber-jwt-crud/models"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	// Hash Password
	if data["password"] != data["password_confirm"] {
		logger.Error("Password do not match", nil)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Password do not match"})

	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)

	// Convert age from string to int
	age, err := strconv.Atoi(data["age"])
	if err != nil {
		logger.Error("Invalid age format", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid age format"})
	}
	user := models.User{
		Name:    data["name"],
		Phone:   data["phone"],
		Address: data["address"],
		Gender:  data["gender"],
		Email:   data["email"],
		// age from request
		Age:      age,
		Password: string(hashedPassword),
	}
	if err := database.DB.Where("email = ?", data["email"]).First(&user).Error; err == nil {
		logger.Error("User already exists", nil)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User already exists"})
	}
	database.DB.Create(&user)
	if err := database.DB.Error; err != nil {
		logger.Error("Failed to create user", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

	}
	logger.Success("User created successfully")
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		logger.Error("Failed to parse JSON", err)
		return err
	}
	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)
	if err := user.Email; err == "" {
		logger.Error("User not found", nil)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	// Validate Password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		logger.Error("Incorrect password", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Incorrect password"})
	}
	// Generate JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"Nme": user.Name,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	t, err := token.SignedString(middleware.SecretKey)
	if err != nil {
		logger.Error("Failed to generate JWT", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	logger.Success("User logged in successfully " + user.Name)

	return c.JSON(fiber.Map{"token": t})
}
