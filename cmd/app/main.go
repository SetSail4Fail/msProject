package main

import (
	"msProject/internal"

	_ "github.com/lib/pq"
)

func main() {
	internal.CreateTable()
}
