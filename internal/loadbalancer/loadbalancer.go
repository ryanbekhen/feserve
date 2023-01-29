package loadbalancer

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func New(path string, servers []string) fiber.Handler {
	balancer := NewRoundRobin(servers)
	return func(c *fiber.Ctx) error {
		host := balancer.Get()
		requestPath := strings.Replace(c.Path(), path, "", -1)
		query := string(c.Request().URI().QueryString())
		if query != "" {
			query = "?" + query
		}
		c.Request().Header.Add("X-Real-IP", c.IP())
		return proxy.Do(c, host+requestPath+query)
	}
}
