package createaccount

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"log"
	"msProject/config"
)

func createAccont(configPath string) {

	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	inputReader := bufio.NewReader(os.Stdin)
	insertQuery := "INSERT INTO accounts (name, password) VALUES ($1, $2)"

	fmt.Println("Enter account Name:")
	name, _ := inputReader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("Enter account Password")
	password, _ := inputReader.ReadString('\n')
	password = strings.TrimSpace(password)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=mydatabase sslmode=disable",
	cfg.DB.Host, cfg.DB.Port, cfg.DB.Name, cfg.DB.Password)
	db, err := sql.Open("postgres", psqlInfo)
	_, err = db.Exec(insertQuery, name, password)
	if err != nil {
		fmt.Printf("Error while creating account.")
		return
	}
}
