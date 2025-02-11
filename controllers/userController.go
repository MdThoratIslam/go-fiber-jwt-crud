package controllers

// 👉 এই লাইনটি নির্দেশ করে যে এই ফাইলটি controllers প্যাকেজের মধ্যে আছে।
import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-fiber-jwt-crud/database"
	logger "go-fiber-jwt-crud/log"
	"go-fiber-jwt-crud/models"
	"os"
	"time"
)

/*
👉 এই লাইনগুলোতে প্রয়োজনীয় প্যাকেজ ও মডিউল ইমপোর্ট করা হয়েছে:
✅ fmt → টেক্সট ফরম্যাট করার জন্য
✅ github.com/gofiber/fiber/v2 → Fiber Framework ব্যবহার করা হয়েছে
✅ go-fiber-jwt-crud/database → ডাটাবেস সংযোগ (DB Connection)
✅ logger "go-fiber-jwt-crud/log" → লগ ম্যানেজমেন্ট সিস্টেম
✅ go-fiber-jwt-crud/models → মডেলস (User Model)
✅ os → ফাইল সিস্টেম অ্যাক্সেস
✅ time → তারিখ ও সময় ব্যবস্থাপনা
*/

// Define a struct for response excluding password
type UserResponse struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Gender  string `json:"gender"`
	Email   string `json:"email"`
	Age     int    `json:"age"`
}

/*
👉 এই struct (স্ট্রাকচার) ইউজারের তথ্য সংরক্ষণ করে কিন্তু পাসওয়ার্ড রাখে না।
✅ এতে নাম, ফোন নম্বর, ঠিকানা, লিঙ্গ, ইমেইল, এবং বয়স থাকবে।
✅ এটি JSON আউটপুট ফরম্যাট সেট করে।
*/
func GetUsers(c *fiber.Ctx) error {
	//👉 GetUsers ফাংশন Fiber Context (c *fiber.Ctx) নেয় এবং একটি ইউজার লিস্ট পাঠায়।
	/*
		📌 c *fiber.Ctx কীভাবে কাজ করে?
		✅ fiber.Ctx হচ্ছে একটি কনটেক্সট অবজেক্ট যা HTTP রিকোয়েস্ট এবং রেসপন্স সংক্রান্ত তথ্য ধারণ করে।
		✅ এটি Go Fiber ফ্রেমওয়ার্কের একটি বিল্ট-ইন ফিচার, যা প্রতিটি API রিকোয়েস্ট হ্যান্ডল করার সময় সংলগ্ন ডাটা ও মেথডস প্রদান করে।
		📌 এটি কী কাজে লাগে?
		1️⃣ রিকোয়েস্টের তথ্য পাওয়া:
		    URL প্যারামিটার (params) নেওয়া
		    কুয়েরি স্ট্রিং (query parameters) পড়া
		    বডি ডাটা (Body data) পড়া
		2️⃣ রেসপন্স পাঠানো:
		    JSON রেসপন্স পাঠানো
		    টেক্সট বা HTML পাঠানো
		    স্ট্যাটাস কোড সেট করা

		3️⃣ মিডলওয়্যার ও অথেনটিকেশন:
		    JWT টোকেন ভেরিফিকেশন
		    রিকোয়েস্ট লগ করা
		    রুট প্রোটেকশন
	*/
	var users []models.User
	// Fetch only selected columns
	result := database.DB.Select("id", "name", "email", "phone", "address", "age", "gender").Find(&users)

	/*
		✅ database.DB.Select(...) → শুধুমাত্র নির্দিষ্ট কলামগুলো নির্বাচন করে
		✅ Find(&users) → সকল ইউজার খুঁজে বের করে এবং users অ্যারের মধ্যে সংরক্ষণ করে।
	*/

	// Check if no users found
	if len(users) == 0 {
		logger.Warn("No users found")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   "No users found",
			"message": "No users available in the database",
		})
	}

	//✅ যদি কোনো ইউজার না পাওয়া যায়, তাহলে "কNo users found" মেসেজ দেখাবে।

	// Check for database error
	if result.Error != nil {
		logger.Error("Failed to fetch users", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to fetch users",
			"details": result.Error.Error(),
		})
	}
	//✅ যদি ডাটাবেস ত্রুটি (error) থাকে, তাহলে ভুল বার্তা রিটার্ন করবে।

	// Map users to response struct to exclude password
	var userResponses []UserResponse
	for _, user := range users {
		userResponses = append(userResponses, UserResponse{
			Name:    user.Name,
			Phone:   user.Phone,
			Address: user.Address,
			Gender:  user.Gender,
			Email:   user.Email,
			Age:     user.Age,
		})
	}
	//✅ পাসওয়ার্ড বাদ দিয়ে ইউজারের তথ্য UserResponse স্ট্রাকচারে সংরক্ষণ করা হয়।

	logger.Success("Users fetched successfully")
	return c.JSON(userResponses)
}
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	database.DB.First(&user, id)
	/*
		✅ c.Params("id") → URL থেকে ID সংগ্রহ করে
		✅ database.DB.First(&user, id) → ডাটাবেস থেকে নির্দিষ্ট ইউজার খোঁজে।
	*/
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
	//✅ যদি ক্লায়েন্ট থেকে পাঠানো JSON ডাটা সমস্যা করে, তাহলে ত্রুটি পাঠানো হয়।

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
	//✅ যদি তারিখ নির্দিষ্ট না করা হয়, তাহলে আজকের তারিখ ব্যবহার করা হবে।

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

	// ✅ লগ ফাইল থেকে ডাটা রিড
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
