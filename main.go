package main

import (
	"Go-Fiber-JWT/routes"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func main() {
	app := fiber.New()

	routes.Routes(app)

	app.Listen(":3000")
}
