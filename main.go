package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/ClientsSharedBill/src/router/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	PORT, erro := strconv.Atoi(os.Getenv("API_PORT"))
	r := routes.SetupRoutes()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if erro != nil {
		log.Fatal("PORT to running undefined")
	}

	fmt.Printf("Listening on port %d", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), r))
}
