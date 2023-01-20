package router

import (
	"path"
	"sort"

	"github.com/gofiber/fiber/v2"
	"github.com/ryanbekhen/feserve/internal/config"
	"github.com/ryanbekhen/feserve/internal/handler"
	"github.com/ryanbekhen/feserve/internal/loadbalancer"
)

func Builder(app *fiber.App) {
	cfg := config.Load()

	app.Get("/ping", handler.PingHandler)

	sort.Slice(cfg.Routes, func(i, j int) bool {
		return cfg.Routes[i].Path > cfg.Routes[j].Path
	})

	for _, route := range cfg.Routes {
		if len(route.Balancer) == 0 {
			app.Get(route.Path, handler.StaticHandler(cfg.PublicDir, route.File))
		} else {
			app.All(path.Join(route.Path, "*"), loadbalancer.New(route.Path, route.Balancer))
		}
	}
}
