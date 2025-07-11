package authjwt

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func AuthHeaderMiddleware(c *fiber.Ctx) error {

	authHeader := c.Get("Authorization")
	token, err := jwt.Parse(authHeader, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}
	claims := token.Claims.(jwt.MapClaims)
	fmt.Print(claims)
	return c.Next()
}

func AuthCookiesMiddleware(c *fiber.Ctx) error {

	cookie := c.Cookies("token")
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("No token cookie found")
	}

	token, err := jwt.Parse(cookie, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}
	claims := token.Claims.(jwt.MapClaims)
	fmt.Print(claims)
	return c.Next()

}
