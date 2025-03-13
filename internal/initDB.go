package internal

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 6444
	user     = "xenous"
	password = "xenous"
	dbname   = "mydatabase"
)

func CreateTable() {
	// Подключение к PostgreSQL (к базе по умолчанию "postgres")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=postgres sslmode=disable",
		host, port, user, password)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Unable to connect to PostgreSQL: %v\n", err)
	}
	defer db.Close()

	// Подключение к новой базе данных
	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Unable to connect to the new database: %v\n", err)
	}
	defer db.Close()

	// Создание таблицы (если она не существует)
	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			);
	`)
	if err != nil {
		log.Fatalf("Unable to create table: %v\n", err)
	}

	log.Println("Table created successfully!")

	// Вставка тестовых данных
	insertUser(db, "HUI BASHOI", "bashoi@example.com")
}

func insertUser(db *sql.DB, name, email string) {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`

	var id int
	err := db.QueryRow(query, name, email).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to insert user: %v\n", err)
	}

	log.Printf("Inserted user with ID: %d\n", id)
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
