package route

import (
	"goblog/handler"

	"github.com/gofiber/fiber/v2"
)

func CommentRoute(app *fiber.App, commentHandler *handler.CommentHandler) {
	app.Post("/comment", commentHandler.Create)
	app.Get("/commentall", commentHandler.GetAll)
	app.Get("/comment/:id", commentHandler.GetsByPostID)
	app.Put("/comment/:id", commentHandler.Update)
	app.Delete("/comment/:id", commentHandler.Delete)
}
