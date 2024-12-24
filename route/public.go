package route

import (
	"goblog/handler"

	"github.com/gofiber/fiber/v2"
)

func PublicRoute(app *fiber.App, postHandler *handler.PostHandler) {
	app.Get("/post", postHandler.GetAllPosts)
	app.Get("/post/:id", postHandler.GetBlogByID)
}
