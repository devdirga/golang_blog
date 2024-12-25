package handler

import (
	"goblog/model"
	"goblog/service"
	"goblog/util"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type CommentHandler struct {
	commentService service.CommentService
}

func NewCommentHandler(commentService service.CommentService) *CommentHandler {
	return &CommentHandler{commentService: commentService}
}

func (h *CommentHandler) Create(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID, _ := strconv.Atoi(claims["id"].(string))
	var comment model.Comment
	if err := c.BodyParser(&comment); err != nil {
		return util.HandleError(c, fiber.StatusBadRequest, "invalid input")
	}
	comment.UserID = userID
	createdPost, err := h.commentService.Create(c.Context(), comment)
	if err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(createdPost)
}

func (h *CommentHandler) GetAll(c *fiber.Ctx) error {
	posts, err := h.commentService.GetAll(c.Context())
	if err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(posts)
}

func (h *CommentHandler) GetsByPostID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return util.HandleError(c, fiber.StatusBadRequest, "invalid id")
	}
	post, err := h.commentService.GetsByPostID(c.Context(), id)
	if err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(post)
}

func (h *CommentHandler) Update(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID, _ := strconv.Atoi(claims["id"].(string))
	ID, err := c.ParamsInt("id")
	if err != nil {
		return util.HandleError(c, fiber.StatusBadRequest, "invalid id")
	}
	var comment model.Comment
	if err := c.BodyParser(&comment); err != nil {
		return util.HandleError(c, fiber.StatusBadRequest, "invalid body")
	}
	comment.ID = ID
	comment.UserID = userID
	data, err := h.commentService.Update(c.Context(), comment)
	if err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(data)
}

func (h *CommentHandler) Delete(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID, _ := strconv.Atoi(claims["id"].(string))
	id, err := c.ParamsInt("id")
	if err != nil {
		return util.HandleError(c, fiber.StatusBadRequest, "invalid id")
	}
	var comment model.Comment
	if err := c.BodyParser(&comment); err != nil {
		return util.HandleError(c, fiber.StatusBadRequest, "invalid body")
	}
	comment.ID = id
	comment.UserID = userID
	err = h.commentService.Delete(c.Context(), comment)
	if err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(fiber.StatusNoContent)
}
