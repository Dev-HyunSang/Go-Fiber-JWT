package main

import (
	"Go-Fiber-JWT/cmd"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	cmd.Router(app)

	app.Listen(":3000")
}
