package handler

import "github.com/gofiber/fiber/v2"

func PingHandler(c *fiber.Ctx) error {
	return c.SendString("pong")
}
