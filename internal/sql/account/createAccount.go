package account

import (
	"database/sql"
	"fmt"
	"log"
	"msProject/config"

	"github.com/google/uuid"
)

type CreateAccRequest struct {
	name     string
	password string
	email    string
}

type AccCfg struct{
	db *sql.DB
}

func (a *AccCfg) CreateAcc(configPath string, CreateAccRequest CreateAccRequest) {
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	
	insertQuery := "INSERT INTO accounts (id, name, password, email) VALUES ($1, $2, $3, $4)"

	id := uuid.New()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=mydatabase sslmode=disable",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Name, cfg.DB.Password)
	db, err := sql.Open("postgres", psqlInfo)

	_, err = db.Exec(insertQuery, id, CreateAccRequest.name, CreateAccRequest.password, CreateAccRequest.email)
	if err != nil {
		log.Printf("Error while creating account.", err)
		return
	}
}

