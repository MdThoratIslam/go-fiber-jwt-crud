package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-fiber-jwt-crud/database"
	"go-fiber-jwt-crud/models"
	"os"
	"time"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	database.DB.First(&user, id)
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	database.DB.First(&user, id)

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Save(&user)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DB.Delete(&models.User{}, id)
	return c.SendString("User Deleted")
}

func GetLog(c *fiber.Ctx) error {
	date := c.Query("date", time.Now().Format("2006-01-02"))
	// ✅ আজকের লগ ফাইলের নাম সেট করা (সঠিক ফরম্যাট)
	fileName := fmt.Sprintf("log/app_%s.log", date)

	// ✅ চেক করুন যে লগ ফাইল আছে কি না
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Println("❌ Log file does not exist:", fileName)
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Could not read log file",
			"details": err.Error(),
		})
	}

	// ✅ লগ ডাটা কনসোলে প্রিন্ট করুন
	fmt.Println(string(data))

	// ✅ JSON রেসপন্স ফেরত দিন
	/*return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Log data printed in console",
		"file":    fileName,
		"data":    string(data),
	})*/
	return c.SendString(string(data))
}
