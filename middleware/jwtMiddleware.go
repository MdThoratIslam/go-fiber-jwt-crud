package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v4"
	"strings"
)

var SecretKey = []byte("your_secret_key")

// ✅ গ্লোবাল Blacklisted Tokens Store (Token Blacklist করার জন্য)
var blacklistedTokens = make(map[string]bool)

// ✅ Middleware: JWT Authentication + Blacklist চেক
func JWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.HS256,
			Key:    SecretKey,
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			//token := c.Get("Authorization")
			//// ✅ Blacklisted Token চেক
			//if blacklistedTokens[token] {
			//	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token is blacklisted"})
			//}
			authHeader := c.Get("Authorization")
			if strings.HasPrefix(authHeader, "Bearer ") {
				token := strings.TrimPrefix(authHeader, "Bearer ")
				if blacklistedTokens[token] {
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token is blacklisted"})
				}
			}
			return c.Next()
		},
	})
}

// ✅ Logout ফাংশন (Blacklist Token)
func Logout(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Token not found"})
	}

	// ✅ টোকেন Blacklist এ যোগ করুন
	blacklistedTokens[token] = true

	return c.JSON(fiber.Map{"message": "Logout successful"})
}
