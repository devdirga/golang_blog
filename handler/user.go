package handler

import (
	"goblog/model"
	"goblog/service"
	"goblog/util"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return util.HandleError(c, fiber.StatusBadRequest, "invalid input")
	}
	createdPost, err := h.userService.Create(c.Context(), user)
	if err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, "cannot create post")
	}
	return c.Status(fiber.StatusCreated).JSON(createdPost)
}

func (h *UserHandler) Signup(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return util.HandleError(c, fiber.StatusBadRequest, "invalid input")
	}
	createdPost, err := h.userService.Signup(c.Context(), user)
	if err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, "cannot create post")
	}
	return c.Status(fiber.StatusCreated).JSON(createdPost)
}

func (h *UserHandler) Signin(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return util.HandleError(c, fiber.StatusBadRequest, "invalid input")
	}
	createdPost, err := h.userService.Signin(c.Context(), user)
	if err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(createdPost)
}

func (h *UserHandler) Google(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return util.HandleError(c, fiber.StatusBadRequest, "invalid input")
	}
	createdPost, err := h.userService.Google(c.Context(), user)
	if err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, "cannot create post")
	}
	return c.Status(fiber.StatusCreated).JSON(createdPost)
}
