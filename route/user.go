package route

import (
	"goblog/handler"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App, userHandler *handler.UserHandler) {
	app.Post("/user", userHandler.Create)
}
