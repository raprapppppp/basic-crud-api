package api

import (
	"go_fiber/db"
	"go_fiber/handlers"
	"go_fiber/repo"
	"go_fiber/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Routes(route fiber.Router) {
	//Database Connection
	db.ConnectionDB()
	//Init
	userRepo := repo.NewUserRepository(db.Database)
	userService := services.UserServiceInit(userRepo)
	userhandler := handlers.NewUserHandler(userService)

	app := route.Group("/users")

	route.Use(cors.New())

	app.Get("/", userhandler.Getuser)
	app.Post("/", userhandler.CreateUser)
	app.Delete("/", userhandler.DeleteUser)
	app.Put("/:id", userhandler.UpdateUser)

	app.Post("/account/create", userhandler.CreateUserAccount)
	app.Post("/account/login", userhandler.LoginUserAccount)
}
