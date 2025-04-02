package operations

import (
	"log"
	"msProject/config"
	"msProject/mypkg/postgres"

	"github.com/google/uuid"
)

type DepositRequest struct {
	Value int
	Uuid  uuid.UUID
}

func deposit(cfg *config.Config, DepositRequest DepositRequest) {
	InsertQuery := "UPDATE accounts SET balance = balance + $1 WHERE id = $2 VALUES ($1, $2)"
	db := postgres.DbConnect(cfg)
 
	_, err := db.Exec(InsertQuery, DepositRequest.Value, DepositRequest.Uuid)
	if err != nil {
		log.Fatalf("User not found", err)
	}
}
