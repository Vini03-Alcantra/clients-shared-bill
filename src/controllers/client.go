package client

import (
	"github.com/ClientsSharedBill/src/database"
	"github.com/ClientsSharedBill/src/models"
	"github.com/gofiber/fiber"
)

func GetClients(c *fiber.Ctx) {
	db := database.Connect()
	var clients []models.Client
	db.Find(clients)
	if len(clients) > 0 {
		c.JSON(nil)
	}
	c.JSON(clients)
}

func GetClient(c *fiber.Ctx) {
	c.Send("book")
}

func NewClient(c *fiber.Ctx) {
	c.Send("new book")
}

func DeleteClient(c *fiber.Ctx) {
	c.Send("delete book")
}
