package app

import (
	"msProject/config"
	"msProject/mypkg/grpc"
	"msProject/mypkg/postgres"

	_ "github.com/lib/pq"
)

func Run(cfg *config.Config) {
	postgres.DbConnect(cfg)
	postgres.CreateTable(cfg)

	grpc.CreateTCP()
}
