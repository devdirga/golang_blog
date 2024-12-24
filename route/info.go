package route

import (
	"goblog/handler"

	"github.com/gofiber/fiber/v2"
)

func InfoRoute(app *fiber.App, userHandler *handler.UserHandler) {
	app.Get("/me", userHandler.Me)
}
