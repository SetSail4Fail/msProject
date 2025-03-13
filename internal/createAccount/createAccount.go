package createaccount

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"github.com/jackc/pgx/v5/pgxpool"
)

func createAccont() {
	inputReader := bufio.NewReader(os.Stdin)
	insertQuery := "INSERT INTO accounts (name, password) VALUES ($1, $2)"

	fmt.Println("Enter account Name:")
	name, _ := inputReader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("Enter account Password")
	password, _ := inputReader.ReadString('\n')
	password = strings.TrimSpace(password)


	connStr := "postgres://xenous:xenous@localhost:6444/accounts"
	db, err := pgxpool.New(context.Background(), connStr)
	_, err = db.Exec(context.Background(), insertQuery, name, password)
	if err != nil {
		fmt.Printf("Error while creating account.")
		return
	}
}
