package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"msProject/config"
)

	var createTableQuery = `
			CREATE TABLE IF NOT EXISTS accounts (
			id uuid PRIMARY KEY,
			name TEXT NOT NULL,
			password TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			);
	`
func CreateTable(cfg *config.Config) {
	db := DbConnect(cfg)
	_, err := db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Unable to create table: %v\n", err)
	}

	log.Println("Table created successfully!")
}

func CfgParse(ConfigPath string) (cfg *config.Config){
	cfg, err := config.NewConfig(ConfigPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	return cfg
}

func DbConnect(cfg *config.Config) (db *sql.DB){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Name, cfg.DB.Password, cfg.DB.DBname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalf("Unable to connect to the new database: %v\n", err)
	}
	return db
}