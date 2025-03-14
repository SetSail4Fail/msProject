package app

import (
	"msProject/mypkg/postgres"

	_ "github.com/lib/pq"
)

func Run(configPath string) {
	postgres.CreateTable(configPath)
}
