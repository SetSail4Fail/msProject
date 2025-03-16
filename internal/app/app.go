package app

import (
	"msProject/mypkg/postgres"
	"msProject/internal/sql/account"

	_ "github.com/lib/pq"
)

func Run(configPath string) {
	postgres.CreateTable(configPath)
	sql.CreateAcc(configPath)
}
