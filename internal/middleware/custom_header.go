package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryanbekhen/feserve/internal/config"
)

func CustomHeaderMiddleware(c *fiber.Ctx) error {
	conf := config.Load()
	c.Set("X-Powered-By", "FESERVE")
	for k, v := range conf.Headers {
		c.Set(k, v)
	}
	return c.Next()
}
