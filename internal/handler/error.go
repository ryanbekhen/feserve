package handler

import (
	"net"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := fiber.ErrInternalServerError.Message

	switch t := err.(type) {
	case *fiber.Error:
		code = t.Code
		message = t.Message
	case *net.OpError:
		if t.Op == "dial" || t.Op == "read" {
			code = fiber.StatusServiceUnavailable
			message = fiber.ErrServiceUnavailable.Message
		}
	}

	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	return c.Status(code).SendString(message)
}
