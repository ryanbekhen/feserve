package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Version      string            `yaml:"version"`
	Host         string            `yaml:"host"`
	Port         string            `yaml:"port"`
	TLSPort      string            `yaml:"tlsPort"`
	Headers      map[string]string `yaml:"headers"`
	AllowOrigins string            `yaml:"allowOrigins"`
	TimeZone     string            `yaml:"timezone"`
	PublicDir    string            `yaml:"publicDir"`
	ProxyHeader  string            `yaml:"proxyHeader"`
	Routes       []Routes          `yaml:"routes"`
	Letsencrypt  *Letsencrypt      `yaml:"letsencrypt,omitempty"`
}

type Routes struct {
	Path     string   `yaml:"path"`
	File     string   `yaml:"file"`
	Rewrite  bool     `yaml:"rewrite"`
	Domain   string   `yaml:"domain"`
	Balancer []string `yaml:"balancer"`
}

type Letsencrypt struct {
	Email     string   `yaml:"email"`
	Domains   []string `yaml:"domains"`
	CertsPath string   `yaml:"certsPath"`
}

const version = "1"

var (
	config     *Config
	configOnce sync.Once
)

func Load() *Config {
	configOnce.Do(func() {
		viper.AddConfigPath(".")
		viper.SetConfigName("app")
		_ = viper.ReadInConfig()
		viper.AutomaticEnv()

		config = &Config{}
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}

		if _, err := os.Stat("app.yaml"); err == nil {
			if config.Version != version {
				panic(fmt.Sprintf("configuration version must be %s\n", version))
			}
		} else {
			config.Version = version
		}

		if viper.GetString("HOST") != "" {
			config.Host = viper.GetString("HOST")
		}

		if viper.GetString("PORT") != "" {
			config.Port = viper.GetString("PORT")
		}

		if viper.GetString("TLS_PORT") != "" {
			config.TLSPort = viper.GetString("TLS_PORT")
		}

		if config.Port == "" && config.Letsencrypt != nil {
			config.Port = "80"
		}

		if config.Port == "" {
			config.Port = "8000"
		}

		if config.TLSPort == "" {
			config.TLSPort = "443"
		}

		if viper.GetString("TZ") != "" {
			config.TimeZone = viper.GetString("TZ")
		}

		if config.TimeZone == "" {
			config.TimeZone = "UTC"
		}

		if config.PublicDir == "" {
			config.PublicDir = "public"
		}

		if len(config.Routes) == 0 {
			config.Routes = append(config.Routes, Routes{
				Path: "*",
				File: "index.html",
			})
		}

		if config.Letsencrypt != nil {
			if config.Letsencrypt.CertsPath == "" {
				config.Letsencrypt.CertsPath = "certs"
			}
		}
	})
	return config
}
