package main

import (
	"log"
	"msProject/config"
	"msProject/internal/app"

	_ "github.com/lib/pq"
)

const ConfigPath = "config/config.yaml"

func main() {
	cfg, err := config.NewConfig(ConfigPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	app.Run(cfg)
}
