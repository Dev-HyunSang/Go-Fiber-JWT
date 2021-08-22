package cmd

import (
	"Go-Fiber-JWT/handler"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// 경로 선언
func Router(app *fiber.App) {
	app.Post("/login", handler.LoginHander)
	app.Get("/", handler.AccessibleHandler)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	app.Get("/restricted", handler.RestrictedHandler)
}
