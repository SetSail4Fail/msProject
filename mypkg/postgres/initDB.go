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

	// Подключение к PostgreSQL (к базе по умолчанию "mydatabase")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=mydatabase sslmode=disable",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Name, cfg.DB.Password)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Unable to connect to PostgreSQL: %v\n", err)
	}
	defer db.Close()

	// Подключение к новой базе данных
	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Name, cfg.DB.Password, cfg.DB.DBname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Unable to connect to the new database: %v\n", err)
	}
	defer db.Close()

	// Создание таблицы (если она не существует)
	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS accounts (
			id uuid PRIMARY KEY,
			name TEXT NOT NULL,
			password TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			);
	`)
	if err != nil {
		log.Fatalf("Unable to create table: %v\n", err)
	}

	log.Println("Table created successfully!")
}


// import (
// 	"context"
// 	"fmt"

// 	"github.com/jackc/pgx/v5/pgxpool"
// )

// func InitDB() {
// 	connStr := "host=localhost port=6444 user=xenous password=xenous dbname=xenous sslmode=disable"
// 	db, err := pgxpool.New(context.Background(), connStr)
// 	if err != nil {
// 		fmt.Println("Connection error.")
// 	} else {
// 		fmt.Println("Connect to Database - SUCCESS.")
// 	}

// 	createTableQuery := `
// 		CREATE TABLE IF NOT EXISTS accounts (
// 			id uuid PRIMARY KEY,
// 			name text NOT NULL,
// 			password text NOT NULL
// 		);
// 	`
// 	_, err = db.Exec(context.Background(), createTableQuery)
// 	if err != nil {
// 		fmt.Printf(err.Error())
// 		return
// 	}
// }
