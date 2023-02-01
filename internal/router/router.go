package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryanbekhen/feserve/internal/config"
	"github.com/ryanbekhen/feserve/internal/handler"
	"github.com/ryanbekhen/feserve/internal/proxy"
)

func Builder(app *fiber.App) {
	cfg := config.Load()
	lb := proxy.New()
	for _, route := range cfg.Routes {
		if len(route.Balancer) == 0 {
			continue
		}

		if route.Domain != "" {
			lb.AddForwardToDomain(route.Domain, route.Balancer)
		} else if route.Path != "" {
			lb.AddForwardToPath(route.Path, route.Rewrite, route.Balancer)
		} else {
			panic("domain or path cannot be empty")
		}

	}
	lb.Routing(app)

	for _, route := range cfg.Routes {
		if len(route.Balancer) > 0 {
			continue
		}
		app.Get(route.Path, handler.StaticHandler(cfg.PublicDir, route.File))
	}
}
