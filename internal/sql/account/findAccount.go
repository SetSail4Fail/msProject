package account

import (
	"log"
	"msProject/config"
	"msProject/mypkg/postgres"

	"github.com/google/uuid"
)

type FindAcc struct {
	Name     string
	Password string
}

type GetUserStr struct{
	Uuid uuid.UUID
}

func (*GetUserStr) FindAccountID(cfg *config.Config, FindAcc FindAcc) (uuid string) {
	db := postgres.DbConnect(cfg)

	query := `
	SELECT id FROM accounts 
	WHERE name = $1 AND password = $2
	LIMIT 1
	`
	err := db.QueryRow(query, FindAcc.Name, FindAcc.Password).Scan(&uuid)
	if err != nil {
		log.Fatalf("User not found", err)
	}
	return uuid
}
