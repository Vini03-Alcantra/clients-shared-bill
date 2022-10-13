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

	// router.POST("/api/v1/clients", func(c *gin.Context) {
	// 	c.JSON(http.StatusCreated, gin.H{
	// 		"message": "Hello world",
	// 	})
	// })

	// routes.SetupRoutes().Run()
	fmt.Printf("Listening on port %d", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), r))
}
