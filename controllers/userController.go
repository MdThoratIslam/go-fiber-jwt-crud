package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-fiber-jwt-crud/database"
	logger "go-fiber-jwt-crud/log"
	"go-fiber-jwt-crud/models"
	"os"
	"time"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	if len(users) == 0 {
		logger.Warn("No users found")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   "No users found",
			"message": "No users available in the database",
		})
	}
	logger.Success("Users fetched successfully")
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	database.DB.First(&user, id)
	if user.ID == 0 {
		logger.Warn("User not found")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   "User not found",
			"message": "User does not exist in the database",
		})
	}
	logger.Success("User fetched successfully")
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	database.DB.First(&user, id)

	if err := c.BodyParser(&user); err != nil {
		logger.Error("Failed to parse JSON", err)
		return err
	}

	database.DB.Save(&user)
	logger.Success("User updated successfully")
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DB.Delete(&models.User{}, id)
	if database.DB.Error != nil {
		logger.Error("Failed to delete user", database.DB.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to delete user",
			"details": database.DB.Error.Error(),
		})
	}
	logger.Success("User deleted successfully with ID: " + id)
	return c.SendString("User Deleted")
}

func GetLog(c *fiber.Ctx) error {
	date := c.Query("date", time.Now().Format("2006-01-02"))
	// ✅ আজকের লগ ফাইলের নাম সেট করা (সঠিক ফরম্যাট)
	fileName := fmt.Sprintf("log/app/app_%s.log", date)
	logger.Success("Log file name set")

	// ✅ চেক করুন যে লগ ফাইল আছে কি না
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Println("❌ Log file does not exist:", fileName)
		logger.Error("Log file not found", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   "Log file not found",
			"file":    fileName,
			"message": "No logs available for today",
		})
	}

	// ✅ লগ ফাইল থেকে ডাটা রিড করুন
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("❌ Could not read log file:", err)
		logger.Error("Could not read log file", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Could not read log file",
			"details": err.Error(),
		})
	}

	// ✅ লগ ডাটা কনসোলে প্রিন্ট করুন
	logger.Success("Log data printed in console")
	// ✅ JSON রেসপন্স ফেরত দিন
	/*return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Log data printed in console",
		"file":    fileName,
		"data":    string(data),
	})*/
	return c.SendString(string(data))
}
