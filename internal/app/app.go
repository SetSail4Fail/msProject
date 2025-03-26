package app

import (
	"msProject/config"
	"msProject/mypkg/grpc"
	"msProject/mypkg/postgres"

	_ "github.com/lib/pq"
)

func Run(cfg *config.Config) {
	CreateDB := postgres.TableCfg{}

	CreateDB.DbConnect(cfg)
	CreateDB.CreateTable(cfg)

	grpc.CreateTCP()
}
