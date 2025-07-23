package main

import (
	"go_fiber/routes/api"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	api.Routes(app)

	log.Fatal(app.Listen(":4000"))

}
