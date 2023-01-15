package router

import (
	"mime"
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
		extension := filepath.Ext(pathFile)

		app.Get(route.Path, func(c *fiber.Ctx) error {
			mimeType := mime.TypeByExtension(extension)
			if mimeType != "" {
				c.Set(fiber.HeaderContentType, mimeType)
			}
			return c.SendFile(pathFile, true)
		})
	}
}
