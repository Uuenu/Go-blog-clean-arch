package main

import (
	"go-blog-ca/internal/app"
	"go-blog-ca/internal/config"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

func main() {

	// Configuration
	cfg := config.GetConfig()

	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)

}
