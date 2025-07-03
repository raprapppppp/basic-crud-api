package main

import (
	"fmt"
	"go_fiber/db"
	"go_fiber/handlers"
	"go_fiber/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {
	dsn := "host=localhost user=postgres password=1234 dbname=employee port=5432 sslmode=disable"
	db.ConnectionDB(dsn)

	app := fiber.New()

	app.Use(basicauth.New(basicauth.Config{
		Realm: "Forbidden",
		Authorizer: func(user, pass string) bool {
			fmt.Print("username", user)
			fmt.Println("password:", pass)

			var login models.Login

			db.Database.Where("username = ?", user).First(&login)

			if user == login.Username && pass == login.Password {
				fmt.Println("eto nga")
				return true
			}

			fmt.Println("b:")
			return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.SendString("Unauthorized")
		},
		ContextUsername: "username",
		ContextPassword: "password",
	}))

	app.Get("/users", handlers.Getuser)
	app.Get("/users/:id", handlers.GetUserById)
	app.Post("/users", handlers.Adduser)
	app.Post("/users/login", handlers.AddLoginUser)
	app.Put("/users/update/:id", handlers.UpdateUser)
	app.Delete("/users/delete/:id", handlers.DeleteUser)
	app.Get("/login/users", handlers.GetLoginUser)

	log.Fatal(app.Listen(":3000"))
}
