package main

import (
	"msProject/internal/app"
	_ "github.com/lib/pq"
)

const configPath = "config/config.yaml"

func main() {
	app.Run(configPath)
}
