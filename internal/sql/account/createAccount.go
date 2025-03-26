package account

import (
	"log"
	"msProject/config"
	"msProject/mypkg/postgres"

	"github.com/google/uuid"
)

type CreateAccRequest struct {
	Name     string
	Password string
	Email    string
}

func CreateAcc(cfg *config.Config, CreateAccRequest CreateAccRequest) {
	InsertQuery := "INSERT INTO accounts (id, name, password, email) VALUES ($1, $2, $3, $4)"
	log.Println("Log of Create acc", CreateAccRequest)
	
	id, err := uuid.NewUUID()
	if err != nil {
		log.Fatalf("Id generation err", err)
	}

	CreateDB := postgres.TableCfg{}
	CreateDB.DbConnect(cfg)

	_, err = CreateDB.GetDB().Exec(InsertQuery, id, CreateAccRequest.Name, CreateAccRequest.Password, CreateAccRequest.Email)
	if err != nil {
		log.Printf("Error while creating account.", err)
		return
	}
}
