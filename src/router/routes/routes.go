package routes

import (
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Get("api/v1/client")
	app.Get("api/v1/client/:id")
	app.Post("api/v1/client")
	app.Delete("api/v1/client/:id")
}
