package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"msProject/config"
)

func CreateTable(configPath string) {

	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Подключение к новой базе данных
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Name, cfg.DB.Password, cfg.DB.DBname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Unable to connect to the new database: %v\n", err)
	}
	defer db.Close()

	// Создание таблицы (если она не существует)
	createTableQuery := `
			CREATE TABLE IF NOT EXISTS accounts (
			id uuid PRIMARY KEY,
			name TEXT NOT NULL,
			password TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			);
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Unable to create table: %v\n", err)
	}

	log.Println("Table created successfully!")
}