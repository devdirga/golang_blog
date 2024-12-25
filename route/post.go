package route

import (
	"goblog/handler"

	"github.com/gofiber/fiber/v2"
)

func PostRoute(app *fiber.App, postHandler *handler.PostHandler) {
	app.Get("/admin/post", postHandler.GetAllAdmin)
	app.Post("/post", postHandler.Create)
	app.Get("/post/:id", postHandler.GetByID)
	app.Put("post/:id", postHandler.Update)
	app.Delete("/post/:id", postHandler.Delete)
}
