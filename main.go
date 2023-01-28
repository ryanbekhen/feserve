package main

import (
	"crypto/tls"
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	mlogger "github.com/gofiber/fiber/v2/middleware/logger"
	certificate "github.com/ryanbekhen/feserve/internal/cert"
	"github.com/ryanbekhen/feserve/internal/config"
	"github.com/ryanbekhen/feserve/internal/handler"
	"github.com/ryanbekhen/feserve/internal/logger"
	mw "github.com/ryanbekhen/feserve/internal/middleware"
	"github.com/ryanbekhen/feserve/internal/router"
)

var conf *config.Config

func init() {
	conf = config.Load()
}

func main() {
	logger := logger.New(logger.Config{
		Timezone: conf.TimeZone,
	})

	if conf.Letsencrypt != nil {
		cert := certificate.NewCert(&certificate.Options{
			Email:     conf.Letsencrypt.Email,
			Domains:   conf.Letsencrypt.Domains,
			CertsPath: conf.Letsencrypt.CertsPath,
			Debug:     true,
		})

		if err := cert.ReadFromFile(); err != nil {
			if err := cert.Generate(); err != nil {
				logger.Error(err.Error())
				os.Exit(1)
			}
		}
	}

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ProxyHeader:           conf.ProxyHeader,
		CompressedFileSuffix:  ".feserve.gz",
		ErrorHandler:          handler.ErrorHandler,
	})

	app.Use(mw.CustomHeaderMiddleware)

	if conf.AllowOrigins != "" {
		app.Use(cors.New(cors.Config{
			AllowOrigins: conf.AllowOrigins,
		}))
	}

	app.Use(cache.New(cache.Config{
		ExpirationGenerator: func(c *fiber.Ctx, cfg *cache.Config) time.Duration {
			cachetime, _ := strconv.Atoi(c.GetRespHeader("Cache-Time", "0"))
			return time.Second * time.Duration(cachetime)
		},
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Path()
		},
	}))

	app.Use(mlogger.New(mlogger.Config{
		TimeZone:   conf.TimeZone,
		TimeFormat: time.RFC3339,
		Format:     "[${time}] - ${ip} - ${status} ${method} ${url} ${ua}\n",
	}))

	app.Static("/", conf.PublicDir, fiber.Static{
		Compress: true,
		ModifyResponse: func(c *fiber.Ctx) error {
			c.Response().Header.Add("Cache-Time", "86400")
			return nil
		},
	})

	router.Builder(app)

	if conf.Letsencrypt != nil {
		certpath := path.Join(conf.Letsencrypt.CertsPath, "ssl.cert")
		keypath := path.Join(conf.Letsencrypt.CertsPath, "ssl.key")

		cer, err := tls.LoadX509KeyPair(certpath, keypath)
		if err != nil {
			logger.Fatal(err)
		}

		config := &tls.Config{
			Certificates: []tls.Certificate{cer},
			MinVersion:   tls.VersionTLS10,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			},
		}

		addr := fmt.Sprintf("%s:%s", conf.Host, "443")

		logger.Info("app listen on ", addr)
		ln, err := tls.Listen("tcp", ":443", config)
		if err != nil {
			panic(err)
		}

		logger.Fatal(app.Listener(ln))
		if err := app.Listener(ln); err != nil {
			logger.Info(err.Error())
			os.Exit(1)
		}
	} else {
		addr := fmt.Sprintf("%s:%s", conf.Host, conf.Port)

		logger.Info("app listen on ", addr)
		if err := app.Listen(addr); err != nil {
			logger.Info(err.Error())
			os.Exit(1)
		}
	}
}
