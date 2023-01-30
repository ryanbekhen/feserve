package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func RedirectHttpsMiddleware(c *fiber.Ctx) error {
	schema := string(c.Request().URI().Scheme())
	if schema != "https" {
		targetUrl := string(c.Context().URI().FullURI())
		targetUrl = strings.Replace(targetUrl, "http://", "https://", 1)
		return c.Redirect(targetUrl, fiber.StatusMovedPermanently)
	}
	return c.Next()
}
