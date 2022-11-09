package config

import (
	"go-blog-ca/pkg/logging"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug"`
	Listen  struct {
		Type   string `yaml:"type"`
		BindIP string `yaml:"bind_ip"`
		Port   string `yaml:"port"`
	} `yaml:"listen"`

	Mongodb struct {
		Host       string `yaml:"host"`
		Port       string `yaml:"port"`
		Database   string `yaml:"database"`
		AuthDb     string `yaml:"authdb"`
		Username   string `yaml:"username"`
		Password   string `yaml:"password"`
		Collection string `yaml:"collection"`
	} `yaml:"mongodb"`

	Session struct {
		TTL            time.Duration `env-required:"true" yaml:"ttl" env:"SESSION_TTL"`
		CookieKey      string        `env-required:"true" yaml:"cookie_key" env:"SESSION_COOKIE_KEY"`
		CookieDomain   string        `yaml:"cookie_domain" env:"SESSION_COOKIE_DOMAIN"`
		CookieSecure   bool          `yaml:"cookie_secure" env:"SESSION_COOKIE_SECURE"`
		CookieHTTPOnly bool          `yaml:"cookie_httponly" env:"SESSION_COOKIE_HTTPONLY"`
	} `yaml:"session"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
