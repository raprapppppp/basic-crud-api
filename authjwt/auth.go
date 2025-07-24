package authjwt

import (
	"fmt"
	"go_fiber/config"
	"slices"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

var errorMessage = "Invalid token"

func AuthHeaderMiddleware(c *fiber.Ctx) error {

	authHeader := c.Get("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).SendString("Missing or invalid Authorization header")
	}

	extractedToken := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(extractedToken, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": errorMessage,
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
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid signing method")
		}
		return []byte(config.Config("COOKIES_SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":err,
		})
	}
	claims := token.Claims.(jwt.MapClaims)
	
	//Get Data on claims
	id := claims["id"].((float64))
	name := claims["username"].(string)
	userRole := claims["role"].(string)

	c.Locals("id", id)
	c.Locals("username", name)
	c.Locals("role", userRole)
	return c.Next()

}

//Role
func RoleBasedMiddleware(requiredRole ...string) fiber.Handler {
	return func (f *fiber.Ctx)error{
		cookie := f.Cookies("token")
		if cookie == "" {
		return f.Status(fiber.StatusUnauthorized).SendString("No token cookie found")
	}
	//Verify and Validate token
	token, err := jwt.Parse(cookie, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid signing method")
		}
		return []byte(config.Config("COOKIES_SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		return f.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err,
		})
	}

	claims := token.Claims.(jwt.MapClaims)
	
	//Get Data on claims
	userRole := claims["role"].(string)
	id := claims["id"].((float64))
	name := claims["username"].(string)


	//Check the user if has the required role to access the route
	if slices.Contains(requiredRole, userRole) {
			f.Locals("username", name)
			f.Locals("id", id)
			f.Locals("role", userRole)
			return f.Next()
		}
		
	return f.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"error" : "Forbidden: you don't have access"})
	}
}

