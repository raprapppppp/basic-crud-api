package main

import (
	"go_fiber/db"
	"go_fiber/handlers"
	"go_fiber/repo"
	"go_fiber/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	dsn := "host=localhost user=postgres password=1234 dbname=employee port=5432 sslmode=disable"
	db.ConnectionDB(dsn)

	//Init
	userRepo := repo.NewUserRepository(db.Database)
	userService := services.UserServiceInit(userRepo)
	userhandler := handlers.NewUserHandler(userService)

	app := fiber.New()

	app.Get("/users", userhandler.Getuser)
	// app.Get("/users/:id", handlers.GetUserById)
	app.Post("/users", userhandler.CreateUser)
	app.Put("/users/update/:id", userhandler.UpdateUser)
	app.Delete("/users/delete/:id", userhandler.DeleteUser)

	log.Fatal(app.Listen(":3000"))
}
