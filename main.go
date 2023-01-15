package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/ryanbekhen/feserve/internal/config"
	mw "github.com/ryanbekhen/feserve/internal/middleware"
	"github.com/ryanbekhen/feserve/internal/router"
)

var conf *config.Config

func init() {
	conf = config.Load()
}

func main() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ProxyHeader:           conf.ProxyHeader,
	})

	app.Use(mw.CustomHeaderMiddleware)
	app.Use(logger.New(logger.Config{
		TimeZone:   conf.TimeZone,
		TimeFormat: time.RFC3339,
		Format:     "[${time}] - ${ip} - ${status} ${method} ${url} ${ua}\n",
	}))

	app.Static("/", conf.PublicDir, fiber.Static{
		Compress: true,
	})

	router.Builder(app)

	addr := fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	loc, err := time.LoadLocation(conf.TimeZone)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	startupMessage := fmt.Sprintf("[%s] - app listening on %s", time.Now().In(loc).Format(time.RFC3339), addr)
	fmt.Println(startupMessage)
	if err := app.Listen(addr); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
