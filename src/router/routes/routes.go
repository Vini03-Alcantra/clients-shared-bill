package routes

import (
	client "github.com/ClientsSharedBill/src/controllers"
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Get("api/v1/clients", client.GetClients)
	app.Get("api/v1/clients/:id", client.GetClient)
	app.Post("api/v1/clients", client.NewClient)
	app.Delete("api/v1/clients/:id", client.DeleteClient)
}
