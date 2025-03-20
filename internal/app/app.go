package app

import (
	"msProject/mypkg/postgres"
	"msProject/mypkg/grpc"

	_ "github.com/lib/pq"
)

func Run(configPath string) {
	postgres.CreateTable(configPath)
	grpc.CreateTCP()
}
