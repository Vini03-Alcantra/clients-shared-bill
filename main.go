package main

import (
	"log"
	"os"
	"strconv"

	"github.com/ClientsSharedBill/src/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	database.Connect()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT, erro := strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		PORT = 3032
	}
	log.Println(PORT)
	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT)))
}
