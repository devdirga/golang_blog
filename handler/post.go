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

func (h *PostHandler) Create(c *fiber.Ctx) error {
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

func (h *PostHandler) GetAll(c *fiber.Ctx) error {
	posts, err := h.postService.GetAll(c.Context())
	if err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(posts)
}

func (h *PostHandler) GetAllAdmin(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authorID, _ := strconv.Atoi(claims["id"].(string))
	posts, err := h.postService.GetAllAdmin(c.Context(), authorID)
	if err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(posts)
}

func (h *PostHandler) GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return util.HandleError(c, fiber.StatusBadRequest, "invalid id")
	}
	post, err := h.postService.GetByID(c.Context(), id)
	if err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(post)
}

func (h *PostHandler) Update(c *fiber.Ctx) error {
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
	updatedBlog, err := h.postService.Update(c.Context(), post)
	if err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(updatedBlog)
}

func (h *PostHandler) Delete(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authorID, _ := strconv.Atoi(claims["id"].(string))
	id, err := c.ParamsInt("id")
	if err != nil {
		return util.HandleError(c, fiber.StatusBadRequest, "invalid id")
	}
	if err := h.postService.Delete(c.Context(), id, authorID); err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}
