package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go-fiber-jwt-crud/BaseResponceForApi"
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

	// Check if passwords match
	if data["password"] != data["password_confirm"] {
		logger.Error("Passwords do not match", nil)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Passwords do not match"})
	}

	// Hash Password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)

	// Convert age from string to int
	age, err := strconv.Atoi(data["age"])
	if err != nil {
		logger.Error("Invalid age format", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid age format"})
	}

	// Check if user already exists
	var existingUser models.User
	if err := database.DB.Where("email = ?", data["email"]).First(&existingUser).Error; err == nil {
		logger.Error("User already exists", nil)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User already exists"})
	}

	// Create new user
	user := models.User{
		Name:     data["name"],
		Phone:    data["phone"],
		Address:  data["address"],
		Gender:   data["gender"],
		Email:    data["email"],
		Age:      age,
		Password: string(hashedPassword),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		logger.Error("Failed to create user", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	logger.Success("User created successfully " + user.Name)
	//response := fiber.Map{
	//	"message": "User created successfully",
	//	"Name":    user.Name,
	//	"Email":   user.Email,
	//	"Phone":   user.Phone,
	//	"Address": user.Address,
	//	"Gender":  user.Gender,
	//	"Age":     user.Age,
	//}
	//return c.JSON(response)

	apiResponse := BaseResponceForApi.ApiResponse{
		Message: "User logged in successfully",
		Status:  "success",
		Data: map[string]interface{}{
			"Name":   user.Name,
			"Email":  user.Email,
			"Phone":  user.Phone,
			"Age":    user.Age,
			"Gender": user.Gender,
		},
	}
	return c.JSON(apiResponse)
}

func Login(c *fiber.Ctx) error {
	// Parse JSON request
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		logger.Error("Failed to parse JSON", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Find user in database
	var user models.User
	if err := database.DB.Where("email = ?", data["email"]).First(&user).Error; err != nil {
		logger.Error("User not found", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Validate password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		fmt.Println(err)
		logger.Error("Incorrect password", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Incorrect password"})
	}

	// Generate JWT Token with 2-minute expiration

	//expirationTime := time.Now().Add(time.Hour * 2) // 2 hours
	expirationTime := time.Now().Add(time.Minute * 2) // 2 minutes
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"name": user.Name,
		"exp":  expirationTime.Unix(),
	})

	// Sign the token
	t, err := token.SignedString([]byte(middleware.SecretKey))
	if err != nil {
		logger.Error("Failed to generate JWT", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate token"})
	}

	// Convert expiration time to formatted string
	formattedTime := expirationTime.Format("2006-01-02 03:04:05 PM")

	// Construct API response
	apiResponse := BaseResponceForApi.ApiResponse{
		Message: "User logged in successfully",
		Status:  "success",
		Data: map[string]interface{}{
			"token":  t,
			"exp":    formattedTime,
			"name":   user.Name,
			"email":  user.Email,
			"phone":  user.Phone,
			"age":    user.Age,
			"gender": user.Gender,
		},
	}

	logger.Success("User logged in successfully !!!\nToken: " + t + "\nExpiry: " + formattedTime + "\n" + "Name" +
		" : " + user.Name + "\nEmail: " + user.Email + "\nPhone: " + user.Phone + "\nAge: " + strconv.Itoa(user.
		Age) + "\nGender: " + user.Gender)
	return c.JSON(apiResponse)
}
