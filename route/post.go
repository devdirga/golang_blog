package route

import (
	"goblog/handler"

	"github.com/gofiber/fiber/v2"
)

func PostRoute(app *fiber.App, postHandler *handler.PostHandler) {
	app.Get("/admin/post", postHandler.GetAllPostsAdmin)
	app.Post("/post", postHandler.CreatePost)
	app.Get("/post/:id", postHandler.GetBlogByID)
	app.Put("post/:id", postHandler.UpdatePost)
	app.Delete("/post/:id", postHandler.DeleteBlog)
}
