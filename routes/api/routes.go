package api

import (
	"fmt"
	"go_fiber/db"
	"go_fiber/handlers"
	"go_fiber/repo"
	"go_fiber/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/golang-jwt/jwt/v5"
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

func Routes(route fiber.Router) {
	//Database Connection
	db.ConnectionDB()
	//Init
	userRepo := repo.NewUserRepository(db.Database)
	userService := services.UserServiceInit(userRepo)
	userhandler := handlers.NewUserHandler(userService)

	route.Use(cors.New())

	app := route.Group("/api", AuthHeaderMiddleware)

	route.Post("/account/create", userhandler.CreateUserAccount)
	route.Post("/account/login", userhandler.LoginUserAccount)

	app.Get("/", userhandler.Getuser)
	app.Post("/", userhandler.CreateUser)
	app.Delete("/", userhandler.DeleteUser)
	app.Put("/:id", userhandler.UpdateUser)

}
