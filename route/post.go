package route

import (
	"goblog/handler"

	"github.com/gofiber/fiber/v2"
)

func PostRoute(app *fiber.App, postHandler *handler.PostHandler) {
	app.Post("/post", postHandler.CreatePost)
	app.Get("/post", postHandler.GetAllPosts)
	app.Patch("/post/:id", postHandler.GetBlogByID)
	app.Put("post/:id", postHandler.UpdatePost)
	app.Delete("/post/:id", postHandler.DeleteBlog)
}
