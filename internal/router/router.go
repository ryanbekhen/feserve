package router

import (
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/ryanbekhen/feserve/internal/config"
	"github.com/ryanbekhen/feserve/internal/handler"
)

func Builder(app *fiber.App) {
	cfg := config.Load()

	app.Get("/ping", handler.Ping)

	for _, route := range cfg.Routes {
		pathFile := filepath.Join(cfg.PublicDir, route.File)

		app.Get(route.Path, func(c *fiber.Ctx) error {
			return c.SendFile(pathFile)
		})
	}
}
