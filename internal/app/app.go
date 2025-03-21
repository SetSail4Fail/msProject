package app

import (
	"msProject/mypkg/grpc"
	"msProject/mypkg/postgres"

	_ "github.com/lib/pq"
)

func Run(configPath string) {
	CreateDB := postgres.TableCfg{}

	CreateDB.DbConnect(configPath)
	CreateDB.CreateTable(configPath)
	
	grpc.CreateTCP()
}