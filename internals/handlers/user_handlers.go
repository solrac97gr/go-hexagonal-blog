package handlers

import (
	"goHexagonalBlog/internal/core/ports"

	fiber "github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
	userService ports.IUserService
}

var _ ports.IUserHandlers = (*UserHandlers)(nil)

func NewUserHandlers(userService ports.IUserService) *UserHandlers {
	return &UserHandlers{
		userService: userService,
	}
}

func (h *UserHandlers) Login(c *fiber.Ctx) error {
	var email string
	var password string
	//Extract the body and get the email and password
	err := h.userService.Login(email, password)
	if err != nil {
		return err
	}
	return nil
}

func (h *UserHandlers) Register(c *fiber.Ctx) error {
	var email string
	var password string
	var confirmPassword string

	//Extract the body and get the email and password
	err := h.userService.Register(email, password, confirmPassword)
	if err != nil {
		return err
	}
	return nil
}
