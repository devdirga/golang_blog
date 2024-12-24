package handler

import (
	"goblog/model"
	"goblog/service"
	"goblog/util"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type PostHandler struct {
	postService service.PostService
}

func NewPostHandler(postService service.PostService) *PostHandler {
	return &PostHandler{postService: postService}
}

func (h *PostHandler) CreatePost(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authorID, _ := strconv.Atoi(claims["id"].(string))

	var post model.Post
	if err := c.BodyParser(&post); err != nil {
		return util.HandleError(c, fiber.StatusBadRequest, "invalid input")
	}

	post.Author = authorID
	createdPost, err := h.postService.CreatePostWithUpdateUser(c.Context(), post)
	if err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(createdPost)
}

func (h *PostHandler) GetAllPosts(c *fiber.Ctx) error {
	posts, err := h.postService.GetAllPosts(c.Context())
	if err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(posts)
}

func (h *PostHandler) GetAllPostsAdmin(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authorID, _ := strconv.Atoi(claims["id"].(string))

	posts, err := h.postService.GetAllPostsAdmin(c.Context(), authorID)
	if err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(posts)
}

func (h *PostHandler) GetBlogByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return util.HandleError(c, fiber.StatusBadRequest, "invalid id")
	}
	post, err := h.postService.GetPostByID(c.Context(), id)
	if err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(post)
}

func (h *PostHandler) UpdatePost(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authorID, _ := strconv.Atoi(claims["id"].(string))

	id, err := c.ParamsInt("id")
	if err != nil {
		return util.HandleError(c, fiber.StatusBadRequest, "invalid id")
	}
	var post model.Post
	if err := c.BodyParser(&post); err != nil {
		return util.HandleError(c, fiber.StatusBadRequest, "invalid body")
	}
	post.ID = id
	post.Author = authorID

	updatedBlog, err := h.postService.UpdatePost(c.Context(), post)
	if err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(updatedBlog)
}

func (h *PostHandler) DeleteBlog(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authorID, _ := strconv.Atoi(claims["id"].(string))

	id, err := c.ParamsInt("id")
	if err != nil {
		return util.HandleError(c, fiber.StatusBadRequest, "invalid id")
	}
	if err := h.postService.DeletePost(c.Context(), id, authorID); err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}
