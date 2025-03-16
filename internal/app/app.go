package app

import (
	"msProject/mypkg/postgres"
	"msProject/internal/createAccount"

	_ "github.com/lib/pq"
)

func Run(configPath string) {
	postgres.CreateTable(configPath)
	// createaccount.CreateAcc(configPath)
}
