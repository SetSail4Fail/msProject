package app

import (
	"msProject/mypkg/grpc"
	"msProject/mypkg/postgres"

	_ "github.com/lib/pq"
)

func Run(ConfigPath string) {
	CreateDB := postgres.TableCfg{}

	CreateDB.DbConnect(ConfigPath)
	CreateDB.CreateTable(ConfigPath)
	
	grpc.CreateTCP()
}