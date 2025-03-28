package main

import (
	"msProject/internal/app"
	"msProject/mypkg/postgres"

	_ "github.com/lib/pq"
)

const ConfigPath = "config/config.yaml"

func main() {
	cfg := postgres.CfgParse(ConfigPath)
	app.Run(cfg)
}
