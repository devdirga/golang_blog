package handler

import (
	"goblog/model"
	"goblog/service"
	"goblog/util"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Signup(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return util.HandleError(c, fiber.StatusBadRequest, "invalid input")
	}
	err := h.userService.Signup(c.Context(), user)
	if err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.StatusOK)
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
		return util.HandleError(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(createdPost)
}

func (h *UserHandler) Me(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authorID, _ := strconv.Atoi(claims["id"].(string))

	var u model.User
	u.ID = authorID

	us, err := h.userService.Me(c.Context(), u)
	if err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(us)
}

func (h *UserHandler) UpdateProfile(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authorID, _ := strconv.Atoi(claims["id"].(string))

	var u model.User
	if err := c.BodyParser(&u); err != nil {
		return util.HandleError(c, fiber.StatusBadRequest, "invalid input")
	}
	u.ID = authorID
	err := h.userService.UpdateProfile(c.Context(), u)
	if err != nil {
		return util.HandleError(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.StatusOK)
}
