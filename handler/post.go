package handler

import (
	"goblog/model"
	"goblog/service"
	"goblog/util"

	"github.com/gofiber/fiber/v2"
)

type PostHandler struct {
	postService service.PostService
}

func NewPostHandler(postService service.PostService) *PostHandler {
	return &PostHandler{postService: postService}
}

func (h *PostHandler) CreatePost(c *fiber.Ctx) error {
	var post model.Post
	if err := c.BodyParser(&post); err != nil {
		return util.HandleError(c, fiber.StatusBadRequest, "invalid input")
	}
	createdPost, err := h.postService.CreatePostWithUpdateUser(c.Context(), post)
	if err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, "cannot create post")
	}
	return c.Status(fiber.StatusCreated).JSON(createdPost)
}

func (h *PostHandler) GetAllPosts(c *fiber.Ctx) error {
	posts, err := h.postService.GetAllPosts(c.Context())
	if err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, "could not fetch posts")
	}
	return c.JSON(posts)
}

func (h *PostHandler) GetBlogByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return util.HandleError(c, fiber.StatusBadRequest, "invalid id")

	}
	post, err := h.postService.GetPostByID(c.Context(), uint(id))
	if err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, "post not found")
	}
	return c.JSON(post)
}

func (h *PostHandler) UpdatePost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return util.HandleError(c, fiber.StatusBadRequest, "invalid id")
	}
	var post model.Post
	if err := c.BodyParser(&post); err != nil {
		return util.HandleError(c, fiber.StatusBadRequest, "invalid body")
	}
	post.ID = uint(id)
	updatedBlog, err := h.postService.UpdatePost(c.Context(), post)
	if err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, "post not found")
	}
	return c.JSON(updatedBlog)
}

func (h *PostHandler) DeleteBlog(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return util.HandleError(c, fiber.StatusBadRequest, "invalid id")
	}
	if err := h.postService.DeletePost(c.Context(), uint(id)); err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, "post not found")
	}
	return c.SendStatus(fiber.StatusNoContent)
}
