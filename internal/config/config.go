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
	Headers      map[string]string `yaml:"headers"`
	AllowOrigins string            `yaml:"allowOrigins"`
	TimeZone     string            `yaml:"timezone"`
	PublicDir    string            `yaml:"publicDir"`
	ProxyHeader  string            `yaml:"proxyHeader"`
	Routes       []Routes          `yaml:"routes"`
}

type Routes struct {
	Path string `yaml:"path"`
	File string `yaml:"file"`
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
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println(err)
		}
		viper.AutomaticEnv()

		config = &Config{}
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}

		if _, err := os.Stat("app.yaml"); err == nil {
			if config.Version != version {
				fmt.Printf("configuration version must be %s\n", version)
				os.Exit(1)
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

		if config.Port == "" {
			config.Port = "8000"
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
	})
	return config
}
