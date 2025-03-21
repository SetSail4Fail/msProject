package account

import (
	"log"
	"msProject/mypkg/postgres"

	"github.com/google/uuid"
)

type CreateAccRequest struct {
	name     string
	password string
	email    string
}

func CreateAcc(configPath string, CreateAccRequest CreateAccRequest) {
	insertQuery := "INSERT INTO accounts (id, name, password, email) VALUES ($1, $2, $3, $4)"

	id := uuid.New()

	CreateDB := postgres.TableCfg{}
	CreateDB.DbConnect(configPath)

	_, err := CreateDB.GetDB().Exec(insertQuery, id, CreateAccRequest.name, CreateAccRequest.password, CreateAccRequest.email)
	if err != nil {
		log.Printf("Error while creating account.", err)
		return
	}
}
