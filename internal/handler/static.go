package handler

import (
	"mime"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func StaticHandler(dir string, file string) fiber.Handler {
	pathFile := filepath.Join(dir, file)
	return func(c *fiber.Ctx) error {
		ext := filepath.Ext(c.OriginalURL())
		mimetype := mime.TypeByExtension(ext)
		c.Set(fiber.HeaderContentType, mimetype)
		c.Response().Header.Add("Cache-Time", "86400")
		return c.SendFile(pathFile)
	}
}
