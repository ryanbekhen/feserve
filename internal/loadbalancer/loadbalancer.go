package loadbalancer

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func New(path string, server []string) fiber.Handler {
	return proxy.Balancer(proxy.Config{
		Servers: server,
		ModifyRequest: func(c *fiber.Ctx) error {
			requestPath := strings.Replace(c.Path(), path, "", -1)
			c.Request().SetRequestURI(requestPath)
			c.Request().Header.Add("X-Real-IP", c.IP())
			c.Request().Header.Add("X-Real-Path", c.Path())
			return nil
		},
		ModifyResponse: func(c *fiber.Ctx) error {
			c.Response().Header.Del(fiber.HeaderServer)
			c.Request().SetRequestURI(c.Get("X-Real-Path"))
			return nil
		},
	})
}
