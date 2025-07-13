package api

import (
	"go_fiber/authjwt"
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

	route.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000", // Your Next.js frontend
		AllowCredentials: true,
	}))

	route.Post("/account/create", userhandler.CreateUserAccount)
	route.Post("/account/login", userhandler.LoginUserAccount)

	//Input your auth Middleware in this to protect route
	app := route.Group("/api", authjwt.AuthCookiesMiddleware)

	app.Get("/", userhandler.Getuser)
	app.Post("/", userhandler.CreateUser)
	app.Delete("/", userhandler.DeleteUser)
	app.Put("/", userhandler.UpdateUser)

}
