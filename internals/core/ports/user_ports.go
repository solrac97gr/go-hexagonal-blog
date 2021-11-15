package ports

import (
	fiber "github.com/gofiber/fiber/v2"
)

type IUserService interface {
	Login(email string, password string) error
	Register(email string, password string, passwordConfirmation string) error
}

type IUserRepository interface {
	Login(email string, password string) error
	Register(email string, password string) error
}

type IUserHandlers interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type IServer interface {
	Initialize()
}
