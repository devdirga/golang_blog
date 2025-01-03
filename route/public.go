package route

import (
	"goblog/handler"

	"github.com/gofiber/fiber/v2"
)

func PublicRoute(app *fiber.App, postHandler *handler.PostHandler) {
	app.Get("/post", postHandler.GetAll)
	app.Get("/post/:id", postHandler.GetByID)
}
