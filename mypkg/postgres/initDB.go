package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"msProject/config"
)

type TableCfg struct {
	db *sql.DB
}

type CfgStruct struct{
	cfg *config.Config
}

	var createTableQuery = `
			CREATE TABLE IF NOT EXISTS accounts (
			id uuid PRIMARY KEY,
			name TEXT NOT NULL,
			password TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			);
	`
func (d *TableCfg) CreateTable(cfg *config.Config) {

	_, err := d.db.Exec(createTableQuery)
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

func (repo *TableCfg) DbConnect(cfg *config.Config) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Name, cfg.DB.Password, cfg.DB.DBname)

	db, err := sql.Open("postgres", psqlInfo)

	repo.db = db

	if err != nil {
		log.Fatalf("Unable to connect to the new database: %v\n", err)
	}
}

func (acc *TableCfg) GetDB() *sql.DB {
	return acc.db
}