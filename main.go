package main

import (
	"go_fiber/db"
	"go_fiber/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	dsn := "host=localhost user=postgres password=1234 dbname=employee port=5432 sslmode=disable"
	db.ConnectionDB(dsn)

	app := fiber.New()

	app.Get("/users", handlers.Getuser)
	app.Get("/users/:id", handlers.GetUserById)
	app.Post("/users", handlers.Adduser)
	app.Put("/users/update/:id", handlers.UpdateUser)
	app.Delete("/users/delete/:id", handlers.DeleteUser)

	log.Fatal(app.Listen(":3000"))
}
