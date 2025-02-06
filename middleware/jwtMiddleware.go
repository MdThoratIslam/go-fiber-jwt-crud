package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v4"
	logger "go-fiber-jwt-crud/log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var SecretKey = []byte("your_secret_key")
var blacklistedTokens = make(map[string]bool) // ব্ল্যাকলিস্টেড টোকেন সংরক্ষণ করার জন্য ম্যাপ
func JWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.HS256,
			Key:    SecretKey,
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "This route requires a valid token"})
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			tokenStr := c.Get("Authorization")

			// ✅ "Bearer " অংশ বাদ দিয়ে শুধুমাত্র টোকেন সংগ্রহ করা
			tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

			// ✅ যদি টোকেন ব্ল্যাকলিস্ট করা হয় তবে ব্লক করা হবে
			if blacklistedTokens[tokenStr] {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "This token is blacklisted"})
			}

			// ✅ টোকেন পার্স করা (ডিকোড ও যাচাই করা)
			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				return SecretKey, nil
			})
			if err != nil || !token.Valid {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
			}

			// ✅ মেয়াদ শেষ হওয়ার সময় বের করা
			// Get expiration time from token claims
			// Get expiration time from token claims
			expirationUnix := int64(token.Claims.(jwt.MapClaims)["exp"].(float64))
			expTime := time.Unix(expirationUnix, 0)
			// Custom formatted expiration time
			formattedExpTime := expTime.Format("02-Jan-2006 03:04 PM") // Example: "05-Mar-2024 02:30 PM"
			if time.Now().Unix() > expirationUnix {
				blacklistedTokens[tokenStr] = true
				logger.Error("Token Expired in JWTMiddleware Middleware: Expiry Time: ", nil)
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "This token has expired on " + formattedExpTime,
				})
			}

			logger.Success("Token Parsed Successfully in JWTMiddleware Middleware: Expiry Time: " + formattedExpTime)

			// ✅ যদি টোকেন বৈধ হয় তবে পরবর্তী middleware বা হ্যান্ডলারে পাঠানো হবে
			return c.Next()
		},
	})
}

/*
✅ Middleware কীভাবে কাজ করে?

    টোকেন চেক করা হয় Authorization হেডার থেকে।
    Blacklist চেক করা হয়, যদি টোকেন blacklist করা হয় তবে Unauthorized রেসপন্স ফেরত দেওয়া হয়।
    JWT টোকেন যাচাই করা হয়, যদি ভুল হয় তাহলে Invalid Token রিটার্ন করবে।
    Token Expiry চেক করা হয়, মেয়াদ শেষ হলে Unauthorized রেসপন্স পাঠানো হয় এবং টোকেন blacklist করা হয়।
    সবকিছু ঠিক থাকলে c.Next() ব্যবহার করে অনুরোধের পরবর্তী ধাপে পাঠানো হয়।
*/

// ✅ Logout Function: Add Token to Blacklist
func Logout(c *fiber.Ctx) error {
	// ✅ "Authorization" হেডার থেকে টোকেন বের করা
	authHeader := c.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Token not found"})
	}

	// ✅ "Bearer " অংশ বাদ দিয়ে আসল টোকেন বের করা
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	// ✅ টোকেন পার্স করা
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	// ✅ টোকেনের মেয়াদ শেষ হয়েছে কিনা চেক করা
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		expirationTime := int64(claims["exp"].(float64))
		if time.Now().Unix() > expirationTime {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token has already expired"})
		}
	}

	// ✅ টোকেন ব্ল্যাকলিস্ট করা (অর্থাৎ, আর ব্যবহার করা যাবে না)
	blacklistedTokens[tokenStr] = true

	// ✅ লগআউট সফল হলে মেসেজ পাঠানো
	return c.JSON(fiber.Map{"message": "Logout successful"})
}

/*
✅ Logout কীভাবে কাজ করে?

    "Authorization" হেডার থেকে টোকেন সংগ্রহ করা হয়।
    টোকেন যাচাই করা হয়, যদি অবৈধ হয় তবে Invalid Token রিটার্ন করা হয়।
    Expiry চেক করা হয়, যদি মেয়াদোত্তীর্ণ হয় তবে Token has already expired রেসপন্স ফেরত দেয়।
    Blacklist এ টোকেন সংরক্ষণ করা হয়, যাতে সেটি ভবিষ্যতে আর ব্যবহার করা না যায়।
    **সফল Logout হলে "Logout successful" রেসপন্স ফেরত
*/
