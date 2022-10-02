package main

import (
	"fmt"

	"github.com/ClientsSharedBill/src/database"
)

func main() {
	database.Connect()
	fmt.Printf("Listening on Port", 3030)
}
