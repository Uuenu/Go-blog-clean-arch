package main

import (
	"go-blog-ca/config"
	"go-blog-ca/internal/app"
	"log"
)

func main() {

	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)

}
