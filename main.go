package main

import (
	"log"
	"os"
	"strconv"

	"github.com/ClientsSharedBill/src/database"
	"github.com/ClientsSharedBill/src/router/routes"
	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	app := fiber.New()
	database.Connect()
	routes.SetupRoutes(app)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT, erro := strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		PORT = 3032
	}

	app.Listen(PORT)
}
