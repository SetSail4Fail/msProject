package main

import (
	"msProject/internal/app"
	_ "github.com/lib/pq"
)

const ConfigPath = "config/config.yaml"

func main() {
	app.Run(ConfigPath)
}
