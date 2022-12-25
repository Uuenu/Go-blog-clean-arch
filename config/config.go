package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

// type Config struct {
// 	IsDebug *bool `yaml:"is_debug"`

// }

type (
	Config struct {
		Listen   `yaml:"listen"`
		Mongodb  `yaml:"mongodb"`
		Session  `yaml:"session"`
		Postgres `yaml:"postgresql"`
	}

	Listen struct {
		Type   string `yaml:"type"`
		BindIP string `yaml:"bind_ip"`
		Port   string `yaml:"port"`
	}

	Mongodb struct {
		Host       string `yaml:"host"`
		Port       string `yaml:"port"`
		Database   string `yaml:"database"`
		AuthDb     string `yaml:"authdb"`
		Username   string `yaml:"username"`
		Password   string `yaml:"password"`
		Collection string `yaml:"collection"`
	}

	Session struct {
		TTL            time.Duration `env-required:"true" yaml:"ttl" env:"SESSION_TTL"`
		CookieKey      string        `yaml:"cookie_key"`
		CookieDomain   string        `yaml:"cookie_domain"`
		CookieSecure   bool          `yaml:"cookie_secure"`
		CookieHTTPOnly bool          `yaml:"cookie_httponly"`
	}

	Postgres struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Database string `yaml:"database"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// var instance *Config
// var once sync.Once

// func GetConfig() *Config {
// 	once.Do(func() {
// 		logger := logging.GetLogger()
// 		logger.Info("read application configuration")
// 		instance = &Config{}
// 		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
// 			help, _ := cleanenv.GetDescription(instance, nil)
// 			logger.Info(help)
// 			logger.Fatal(err)
// 		}
// 	})
// 	return instance
// }
