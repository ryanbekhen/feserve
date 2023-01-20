package handler

import (
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func StaticHandler(dir string, file string) fiber.Handler {
	pathFile := filepath.Join(dir, file)
	return func(c *fiber.Ctx) error {
		c.Response().Header.Add("Cache-Time", "86400")
		return c.SendFile(pathFile)
	}
}
