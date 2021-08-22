package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func LoginHander(c *fiber.Ctx) error {
	user := c.FormValue("user")
	password := c.FormValue("password")

	if user != "hyunsang" || password != "park" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user + password
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func AccessibleHandler(c *fiber.Ctx) error {
	return c.SendString("Accessible")
}

func RestrictedHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	clamis := user.Claims.(jwt.MapClaims)
	name := clamis["name"].(string)
	return c.SendString("Welecome " + name)
}
