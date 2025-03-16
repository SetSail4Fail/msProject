package main

import(
	"msProject/mypkg/postgres"
	_ "github.com/lib/pq"
)

const configPath = "config/config.yaml"

func main() {
	postgres.CreateTable(configPath)
}
