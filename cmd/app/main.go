package main

import (
	"go-blog-ca/internal/app"
	"go-blog-ca/internal/config"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

func main() {

	cfg := config.GetConfig()

	err := cleanenv.ReadConfig("./config/local.yml", &cfg)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)

}
