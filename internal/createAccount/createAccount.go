package createaccount

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"msProject/config"
	"os"
	"strings"

	"github.com/google/uuid"
)

func CreateAcc(configPath string) {

	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	inputReader := bufio.NewReader(os.Stdin)
	insertQuery := "INSERT INTO accounts (id, name, password, email) VALUES ($1, $2, $3, $4)"

	fmt.Println("Enter account Name:")
	name, _ := inputReader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("Enter account Password")
	password, _ := inputReader.ReadString('\n')
	password = strings.TrimSpace(password)

	fmt.Println("Enter email")
	email, _ := inputReader.ReadString('\n')
	email = strings.TrimSpace(email)

	id := uuid.New()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=mydatabase sslmode=disable",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Name, cfg.DB.Password)
	db, err := sql.Open("postgres", psqlInfo)

	_, err = db.Exec(insertQuery, id, name, password, email)
	if err != nil {
		fmt.Printf("Error while creating account.", err)
		return
	}
}
