package controllers

import (
	"log"
	"net/http"

	"github.com/ClientsSharedBill/src/database"
	"github.com/ClientsSharedBill/src/models"
	"github.com/gin-gonic/gin"
)

func GetCloentsAll(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Hello World",
	})
}

func GetClients(c *gin.Context) {
	var clients []models.Client
	db, connErr := database.GetDatabaseConnection()

	if connErr != nil {
		log.Panic(connErr)
		c.JSON(200, gin.H{
			"message": "Service is unavailable",
		})
		return
	}
	db.Find(clients)
	if len(clients) > 0 {
		c.JSON(200, gin.H{
			"message": clients,
		})
	}
	c.JSON(200, gin.H{
		"data": clients,
	})
}

func GetClient(c *gin.Context) {
	c.JSON(200, "book")
}

func PostClient(c *gin.Context) {
	c.JSON(200, "new book")
}

func DeleteClient(c *gin.Context) {
	c.JSON(200, "delete book")
}
