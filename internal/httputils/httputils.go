package httputils

import "github.com/gofiber/fiber/v2"

func ForwardUserIP(c *fiber.Ctx) {
	c.Request().Header.Add("X-Real-IP", c.IP())
}
