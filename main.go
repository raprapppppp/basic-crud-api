package main

import (
	"go_fiber/db"
	"go_fiber/handlers"
	"go_fiber/repo"
	"go_fiber/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	dsn := "host=localhost user=postgres password=1234 dbname=employee port=5432 sslmode=disable"
	db.ConnectionDB(dsn)

	//Init
	userRepo := repo.NewUserRepository(db.Database)
	userService := services.UserServiceInit(userRepo)
	userhandler := handlers.NewUserHandler(userService)

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/users", userhandler.Getuser)
	app.Post("/users", userhandler.CreateUser)
	app.Delete("/users/", userhandler.DeleteUser)
	app.Put("/users/:id", userhandler.UpdateUser)

	log.Fatal(app.Listen(":3001"))
}
