package main

import (
	"crypto/tls"
	"fmt"
	"os"
	"os/signal"
	"path"
	"strconv"
	"syscall"
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
			Debug:     false,
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
	app.Use(mw.RedirectHttpsMiddleware)

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

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	if conf.Letsencrypt != nil {
		go func() {
			certpath := path.Join(conf.Letsencrypt.CertsPath, "ssl.cert")
			keypath := path.Join(conf.Letsencrypt.CertsPath, "ssl.key")

			cer, err := tls.LoadX509KeyPair(certpath, keypath)
			if err != nil {
				logger.Error(err)
				os.Exit(1)
			}

			tlsConfig := &tls.Config{
				Certificates: []tls.Certificate{cer},
				MinVersion:   tls.VersionTLS12,
				CipherSuites: []uint16{
					tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
					tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
					tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
					tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				},
			}

			addr := fmt.Sprintf("%s:%s", conf.Host, conf.TLSPort)

			logger.Info("app listen on ", addr)
			ln, err := tls.Listen("tcp", addr, tlsConfig)
			if err != nil {
				logger.Error(err.Error())
				os.Exit(1)
			}

			if err := app.Listener(ln); err != nil {
				logger.Error(err.Error())
				os.Exit(1)
			}
		}()
	}

	go func() {
		addr := fmt.Sprintf("%s:%s", conf.Host, conf.Port)

		logger.Info("app listen on ", addr)
		if err := app.Listen(addr); err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
	}()
	<-stop
}
