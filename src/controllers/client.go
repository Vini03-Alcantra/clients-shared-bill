package client

import "github.com/gofiber/fiber"

func GetClients(c *fiber.Ctx) {
	c.Send("All books")
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
